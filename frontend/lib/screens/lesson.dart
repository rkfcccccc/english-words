import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:english_words/utils/api/vocabulary/models.dart';
import 'package:english_words/widgets/floating_container.dart';
import 'package:english_words/widgets/floating_wrapper.dart';
import 'package:english_words/widgets/word_entry_view.dart';
import 'package:flutter/material.dart';
import 'package:english_words/utils/api/vocabulary/vocabulary.dart'
    as vocabulary_api;
import 'package:sizer/sizer.dart';

const initialPage = 100000000;

class LessonScreen extends StatefulWidget {
  const LessonScreen({Key? key}) : super(key: key);

  @override
  State<LessonScreen> createState() => _LessonScreenState();
}

class _LessonScreenState extends State<LessonScreen> {
  late PageController _pageController;
  late Map<int, Challenge> challenges;

  @override
  void initState() {
    _pageController = PageController(initialPage: initialPage);

    challenges = {};
    vocabulary_api.getChallenge().then((challenge) {
      precacheChallenge(challenge);

      setState(() {
        challenges[initialPage] = challenge;
      });
    });

    super.initState();
  }

  void precacheChallenge(Challenge challenge) {
    final pictures = challenge.entry.pictures;
    if (pictures == null || pictures.isEmpty || challenge.learningStep != 0) {
      return;
    }

    for (int i = 0; i < min(pictures.length * 2 / 3, 1); i++) {
      precacheImage(CachedNetworkImageProvider(pictures[i].url), context);
    }
  }

  Future<void> finishChallenge(String action) async {
    final currentPage = _pageController.page!.toInt();
    final entry = challenges[currentPage]!.entry;

    _pageController.nextPage(
      duration: const Duration(milliseconds: 500),
      curve: Curves.easeOutExpo,
    );

    await vocabulary_api.finishChallenge(action, entry.id);

    if (action == "promote") {
      final nextChallenge = await vocabulary_api.getChallenge();
      precacheChallenge(nextChallenge);

      setState(() {
        challenges[currentPage + 1] = nextChallenge;
      });

      challenges.remove(currentPage);
    } else if (action == "resist") {
      final nextChallenge = challenges[currentPage]!.copyWith(learningStep: 0);
      precacheChallenge(nextChallenge);

      setState(() {
        challenges[currentPage + 1] = nextChallenge;
      });

      challenges.remove(currentPage);
    }
  }

  Future<void> setAlreadyLearned() async {
    final currentPage = _pageController.page!.toInt();
    final entry = challenges[currentPage]!.entry;

    _pageController.nextPage(
      duration: const Duration(milliseconds: 500),
      // curve: Curves.easeInOutQuart,
      curve: Curves.easeOutExpo,
    );

    await vocabulary_api.setAlreadyLearned(entry.id, true);

    final nextChallenge = await vocabulary_api.getChallenge();

    setState(() {
      challenges[currentPage + 1] = nextChallenge;
    });

    challenges.remove(currentPage);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView.builder(
        controller: _pageController,
        scrollDirection: Axis.vertical,
        physics: const NeverScrollableScrollPhysics(),
        itemBuilder: (context, index) {
          if (!challenges.containsKey(index)) {
            return FloatingWrapper(
              child: Padding(
                padding: EdgeInsets.all(4.w),
                child: Row(
                  children: [
                    SizedBox.square(
                      dimension: 3.w,
                      child: const CircularProgressIndicator(
                        strokeWidth: 3,
                        color: Color.fromRGBO(70, 70, 70, 1),
                      ),
                    ),
                    SizedBox(width: 4.w),
                    Flexible(
                      child: Text(
                        "We are choosing the right words for you...",
                        style: TextStyle(
                          fontSize: 13.sp,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            );
          }

          final challenge = challenges[index]!;

          if (challenge.learningStep > 0) {
            return FloatingWrapper(
              actions: FloatingActions(
                onPrimary: () => finishChallenge("promote"),
                primaryText: "Yes",
                onSecondary: () => finishChallenge("resist"),
                secondaryText: "I don't remember",
              ),
              child: FloatingContainer(
                child: Padding(
                  padding:
                      EdgeInsets.symmetric(horizontal: 4.w, vertical: 10.w),
                  child: Column(
                    children: [
                      Text(
                        "Do you remember the meaning of",
                        style: TextStyle(fontSize: 12.sp),
                      ),
                      Text(
                        "${challenge.entry.word}?",
                        style: TextStyle(
                          fontSize: 14.sp,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
          }

          return FloatingWrapper(
            actions: FloatingActions(
              onPrimary: () => finishChallenge("promote"),
              primaryText: "Next",
              onSecondary: () => setAlreadyLearned(),
              secondaryText: "I already know this word",
            ),
            child: WordEntryView(entry: challenge.entry),
          );
        },
      ),
    );
  }
}
