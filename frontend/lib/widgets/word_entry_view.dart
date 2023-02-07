import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class WordEntryView extends StatefulWidget {
  final WordEntry entry;
  final Widget? actions;

  const WordEntryView({
    Key? key,
    this.actions,
    required this.entry,
  }) : super(key: key);

  @override
  State<WordEntryView> createState() => _WordEntryViewState();
}

class _WordEntryViewState extends State<WordEntryView> {
  void precache() async {
    final pictures = widget.entry.pictures;
    if (pictures == null) {
      return;
    }

    for (int i = 0; i < pictures.length ~/ 2; i++) {
      if (!mounted) return;
      await precacheImage(CachedNetworkImageProvider(pictures[i].url), context);
    }
  }

  @override
  void initState() {
    WidgetsBinding.instance.addPostFrameCallback((_) => precache());
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (widget.entry.pictures != null)
          _Pictures(pictures: widget.entry.pictures!),
        Padding(
          padding: EdgeInsets.all(4.w),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        widget.entry.word,
                        style: TextStyle(
                          fontSize: 19.sp,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                      if (widget.entry.phonetic != "")
                        Text(
                          widget.entry.phonetic,
                          style: const TextStyle(
                            fontSize: 14,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                    ],
                  ),
                  ElevatedButton(
                    onPressed: () {},
                    style: ElevatedButton.styleFrom(
                      minimumSize: Size.zero, // Set this
                      padding: EdgeInsets.all(2.w),
                      tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                      backgroundColor: const Color.fromRGBO(44, 44, 44, 1),
                      foregroundColor: const Color.fromRGBO(74, 74, 74, 1),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(7),
                      ),
                    ),
                    child: Icon(
                      Icons.translate_rounded,
                      size: 14.sp,
                      color: const Color.fromRGBO(222, 222, 222, 1),
                    ),
                  ),
                ],
              ),
              if (widget.entry.translations != null) ...[
                SizedBox(height: 1.w),
                Wrap(
                  spacing: 1.w,
                  runSpacing: 1.w,
                  children: widget.entry.translations!
                      .map(
                        (translation) => Container(
                          padding: EdgeInsets.all(1.w),
                          decoration: BoxDecoration(
                            color: Theme.of(context).hoverColor,
                            borderRadius: BorderRadius.circular(7),
                          ),
                          child: Text(
                            translation,
                            style: TextStyle(
                              fontSize: 10.sp,
                              fontWeight: FontWeight.w400,
                              color: const Color.fromRGBO(222, 222, 222, 1),
                            ),
                          ),
                        ),
                      )
                      .toList(),
                ),
              ],
              const Divider(),
              _WordMeanings(meanings: widget.entry.meanings),
              if (widget.actions != null) ...[
                SizedBox(height: 4.w),
                widget.actions!,
              ]
            ],
          ),
        ),
      ],
    );
  }
}

class _Pictures extends StatelessWidget {
  const _Pictures({
    Key? key,
    required this.pictures,
  }) : super(key: key);

  final List<SourcedPicture> pictures;

  @override
  Widget build(BuildContext context) {
    return Container(
      color: const Color.fromRGBO(43, 43, 43, 1),
      height: 25.h,
      child: PageView.builder(
        itemBuilder: (ctx, index) => Stack(
          fit: StackFit.expand,
          children: [
            CachedNetworkImage(
              imageUrl: pictures[index].url,
              fit: BoxFit.cover,
            ),
            Positioned(
              bottom: 1.w,
              left: 1.w,
              child: Container(
                padding: EdgeInsets.all(1.w),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(7),
                  color: const Color.fromRGBO(33, 33, 33, 0.7),
                ),
                child: Text(
                  pictures[index].source,
                  style: TextStyle(fontSize: 8.sp),
                ),
              ),
            ),
          ],
        ),
        itemCount: pictures.length,
      ),
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
            (definitions.length > 1 ? "${index + 1}. " : "") +
                definitions[index].text,
            style: TextStyle(
              fontSize: 11.sp,
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
