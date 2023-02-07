import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:english_words/widgets/word_entry_view.dart';
import 'package:flutter/material.dart';

import 'floating_wrapper.dart';

class WrappedWordEntryView extends StatelessWidget {
  final WordEntry entry;
  final Widget? actions;

  const WrappedWordEntryView({
    super.key,
    required this.entry,
    this.actions,
  });

  @override
  Widget build(BuildContext context) {
    return FloatingWrapper(
      children: [
        WordEntryView(
          entry: entry,
          actions: actions,
        ),
        Column(
          children: [
            Text('Context Examples'),
          ],
        )
      ],
    );
  }
}
