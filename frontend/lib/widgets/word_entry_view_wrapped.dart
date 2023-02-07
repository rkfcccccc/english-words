import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:english_words/widgets/word_entry_view.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

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
        Padding(
          padding: EdgeInsets.all(4.w),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'Context Examples',
                style: TextStyle(
                  fontSize: 16.sp,
                  fontWeight: FontWeight.w600,
                ),
              ),
              SizedBox(height: 2.w),
              Container(
                padding: EdgeInsets.all(3.w),
                decoration: BoxDecoration(
                  color: Theme.of(context).hoverColor,
                  borderRadius: BorderRadius.circular(9),
                ),
                child: Text(
                  'The method includes mixing the ore with carbonaceous reducing agents, roasting the resultant mixture and cooling, grinding and separating same.',
                  style: TextStyle(
                    color: Theme.of(context).textTheme.caption!.color,
                    fontSize: 11.sp,
                    fontWeight: FontWeight.w400,
                  ),
                ),
              ),
              SizedBox(height: 2.w),
              Container(
                padding: EdgeInsets.all(3.w),
                decoration: BoxDecoration(
                  color: Theme.of(context).hoverColor,
                  borderRadius: BorderRadius.circular(9),
                ),
                child: Text(
                  'The method includes mixing the ore with carbonaceous reducing agents, roasting the resultant mixture and cooling, grinding and separating same.',
                  style: TextStyle(
                    color: Theme.of(context).textTheme.caption!.color,
                    fontSize: 11.sp,
                    fontWeight: FontWeight.w400,
                  ),
                ),
              ),
            ],
          ),
        )
      ],
    );
  }
}
