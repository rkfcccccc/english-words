import 'dart:math';

import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class WordEntryView extends StatelessWidget {
  final WordEntry entry;

  const WordEntryView({
    Key? key,
    required this.entry,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Container(
          color: const Color.fromRGBO(43, 43, 43, 1),
          height: 25.h,
          child: PageView.builder(
            itemBuilder: (ctx, index) => Image.network(
              "https://images.unsplash.com/photo-1537420327992-d6e192287183?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8c3BhY2V8ZW58MHx8MHx8&auto=format&fit=crop&w=900&q=60",
              width: double.infinity,
              fit: BoxFit.fitWidth,
            ),
            itemCount: 3,
          ),
        ),
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
              // Row(
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
            (definitions.length > 1 ? "${index + 1}. " : "") +
                definitions[index].text,
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
