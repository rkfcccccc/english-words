import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:english_words/widgets/search_item.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'floating_wrapper.dart';
import 'word_entry_view.dart';

class WordSearchItem extends StatelessWidget {
  final WordEntry entry;

  const WordSearchItem({
    Key? key,
    required this.entry,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SearchItem(
      onPressed: () => Navigator.of(context).push(
        MaterialPageRoute(
          builder: (context) => Scaffold(
            body: FloatingWrapper(
              child: WordEntryView(
                entry: entry,
              ),
            ),
          ),
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "definitions of",
                style: TextStyle(
                  color: const Color.fromRGBO(140, 140, 140, 1),
                  fontSize: 10.sp,
                ),
              ),
              Text(
                entry.word,
                style: TextStyle(
                  color: Colors.white,
                  fontWeight: FontWeight.w500,
                  fontSize: 12.sp,
                ),
              ),
            ],
          ),
          SizedBox(width: 1.w),
          const Icon(Icons.chevron_right_rounded),
        ],
      ),
    );
  }
}
