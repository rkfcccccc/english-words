import 'dart:math';

import 'package:english_words/utils/api/vocabulary/models.dart';
import 'package:english_words/widgets/floating_container.dart';
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
    vocabulary_api.getChallenge().then((value) {
      setState(() {
        challenges[initialPage] = value;
      });
    });

    super.initState();
  }

  Future<void> finishChallenge(String action) async {
    final currentPage = _pageController.page!.toInt();
    final entry = challenges[currentPage]!.entry;

    _pageController.nextPage(
      duration: const Duration(milliseconds: 500),
      // curve: Curves.easeInOutQuart,
      curve: Curves.easeOutExpo,
    );

    await vocabulary_api.finishChallenge(action, entry.id);

    if (action == "promote") {
      final nextChallenge = await vocabulary_api.getChallenge();

      setState(() {
        challenges[currentPage + 1] = nextChallenge;
      });

      challenges.remove(currentPage);
    } else if (action == "resist") {
      final nextChallenge = challenges[currentPage]!.copyWith(learningStep: 0);

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

  // Future<void> () async {
  //   final currentPage = _pageController.page!.toInt();
  //   final entry = challenges[currentPage]!.entry;

  //   _pageController.nextPage(
  //     duration: const Duration(milliseconds: 500),
  //     curve: Curves.easeInOutQuart,
  //   );

  //   await vocabulary_api.finishChallenge("resist", entry.id);

  //   final nextChallenge = await vocabulary_api.getChallenge();
  //   setState(() {
  //     challenges[currentPage + 1] = nextChallenge;
  //   });
  // }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView.builder(
        controller: _pageController,
        scrollDirection: Axis.vertical,
        physics: const NeverScrollableScrollPhysics(),
        itemBuilder: (context, index) {
          if (!challenges.containsKey(index)) {
            return _Wrapper(
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
            return _Wrapper(
              actions: _Actions(
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

          return _Wrapper(
            actions: _Actions(
              onPrimary: () => finishChallenge("promote"),
              primaryText: "Next",
              onSecondary: () => setAlreadyLearned(),
              secondaryText: "I already know this word",
            ),
            child: _WordEntry(entry: challenge.entry),
          );
        },
      ),
    );
  }
}

class _Wrapper extends StatefulWidget {
  final Widget child;
  final Widget? actions;

  const _Wrapper({
    Key? key,
    required this.child,
    this.actions,
  }) : super(key: key);

  @override
  State<_Wrapper> createState() => _WrapperState();
}

class _WrapperState extends State<_Wrapper> {
  late ScrollController _controller;
  bool floating = false;

  void updateFloating() {
    assert(_controller.position.hasContentDimensions);

    if (_controller.position.maxScrollExtent != 0) {
      if (!floating) {
        setState(() {
          floating = true;
        });
      }
    } else {
      if (floating) {
        setState(() {
          floating = false;
        });
      }
    }
  }

  @override
  void initState() {
    _controller = ScrollController();
    super.initState();

    WidgetsBinding.instance.addPostFrameCallback((_) => updateFloating());
  }

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: SingleChildScrollView(
        clipBehavior: Clip.none,
        controller: _controller,
        physics: const AlwaysScrollableScrollPhysics(),
        child: Padding(
          padding: EdgeInsets.all(6.w),
          child: Column(
            children: [
              FloatingContainer(
                child: Column(
                  children: [
                    widget.child,
                    if (widget.actions != null && !floating) ...[
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 4.w),
                        child: const Divider(height: 0),
                      ),
                      widget.actions!,
                    ]
                  ],
                ),
              ),
              if (widget.actions != null && floating) ...[
                SizedBox(height: 6.w),
                FloatingContainer(child: widget.actions!),
              ]
            ],
          ),
        ),
      ),
    );
  }
}

class _Actions extends StatelessWidget {
  final VoidCallback? onPrimary, onSecondary;
  final String? primaryText, secondaryText;

  const _Actions({
    Key? key,
    this.onPrimary,
    this.onSecondary,
    this.primaryText,
    this.secondaryText,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(4.w),
      child: Column(
        children: [
          if (onPrimary != null)
            SizedBox(
              width: double.infinity,
              child: ElevatedButton(
                onPressed: onPrimary,
                child: Text(primaryText!, style: TextStyle(fontSize: 12.sp)),
              ),
            ),
          if (onPrimary != null && onSecondary != null) SizedBox(height: 1.w),
          if (onSecondary != null)
            SizedBox(
              width: double.infinity,
              height: 8.w,
              child: TextButton(
                onPressed: onSecondary,
                child: Text(
                  secondaryText!,
                  style: TextStyle(fontSize: 11.sp),
                ),
              ),
            ),
        ],
      ),
    );
  }
}

class _WordEntry extends StatelessWidget {
  final WordEntry entry;
  const _WordEntry({
    Key? key,
    required this.entry,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        // Image.network(
        //   "https://images.unsplash.com/photo-1537420327992-d6e192287183?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8c3BhY2V8ZW58MHx8MHx8&auto=format&fit=crop&w=900&q=60",
        //   height: 25.h,
        //   width: double.infinity,
        //   fit: BoxFit.fitWidth,
        // ),
        Padding(
          padding: EdgeInsets.all(4.w),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                entry.word,
                style: TextStyle(
                  fontSize: 19.sp,
                  fontWeight: FontWeight.w500,
                ),
              ),
              if (entry.phonetic != "")
                Text(
                  entry.phonetic,
                  style: const TextStyle(
                    fontSize: 14,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              const Divider(),
              _WordMeanings(meanings: entry.meanings),
            ],
          ),
        ),
      ],
    );
  }
}

class _WordMeanings extends StatelessWidget {
  final List<Meaning> meanings;
  const _WordMeanings({
    Key? key,
    required this.meanings,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.separated(
      itemBuilder: (context, index) => Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            meanings[index].partOfSpeech,
            style: TextStyle(
              fontSize: 11.sp,
              fontWeight: FontWeight.w600,
            ),
          ),
          _WordDefinitions(
            definitions: meanings[index].definitions,
          ),
        ],
      ),
      shrinkWrap: true,
      separatorBuilder: (context, index) => const Divider(),
      physics: const NeverScrollableScrollPhysics(),
      itemCount: meanings.length,
    );
  }
}

class _WordDefinitions extends StatelessWidget {
  final List<Definition> definitions;
  const _WordDefinitions({
    Key? key,
    required this.definitions,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.separated(
      itemBuilder: (ctx, index) => Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "${index + 1}. ${definitions[index].text}",
            style: TextStyle(
              fontSize: 11.5.sp,
              fontWeight: FontWeight.w400,
            ),
          ),
          if (definitions[index].example != "")
            _WordExample(example: definitions[index].example),
        ],
      ),
      shrinkWrap: true,
      separatorBuilder: (ctx, index) => const SizedBox(height: 6.0),
      physics: const NeverScrollableScrollPhysics(),
      itemCount:
          min(definitions.length, 3), // TODO: add a button to increase it
    );
  }
}

class _WordExample extends StatelessWidget {
  final String example;
  const _WordExample({
    Key? key,
    required this.example,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const SizedBox(height: 4),
        Container(
          padding: const EdgeInsets.symmetric(
            horizontal: 8.0,
            vertical: 8.0,
          ),
          decoration: BoxDecoration(
            color: Theme.of(context).hoverColor,
            borderRadius: BorderRadius.circular(7),
          ),
          child: Text(
            example,
            style: TextStyle(
              fontSize: 10.sp,
              fontWeight: FontWeight.w400,
              color: Theme.of(context).textTheme.caption!.color,
            ),
          ),
        ),
        const SizedBox(height: 4)
      ],
    );
  }
}
