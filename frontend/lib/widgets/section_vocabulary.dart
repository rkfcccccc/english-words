import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'section.dart';

class VocabularySection extends StatelessWidget {
  const VocabularySection({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Section(
      header: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(
            'Your vocabulary',
            style: TextStyle(
              fontSize: 11.sp,
              fontWeight: FontWeight.w600,
            ),
          ),
          Container(
            padding: EdgeInsets.symmetric(
              vertical: 1.w,
              horizontal: 2.w,
            ),
            decoration: BoxDecoration(
              color: Colors.purple,
              borderRadius: BorderRadius.circular(13),
            ),
            child: const Text("65 words"),
          ),
        ],
      ),
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: const [
              Text(
                "Just starting",
                style: TextStyle(color: Color.fromRGBO(160, 160, 160, 1)),
              ),
              Text(
                "23 words",
                style: TextStyle(color: Color.fromRGBO(100, 100, 100, 1)),
              ),
            ],
          ),
          SizedBox(height: 3.w),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: const [
              Text(
                "Moderate knowledge",
                style: TextStyle(color: Color.fromRGBO(160, 160, 160, 1)),
              ),
              Text(
                "32 words",
                style: TextStyle(color: Color.fromRGBO(100, 100, 100, 1)),
              ),
            ],
          ),
          SizedBox(height: 3.w),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: const [
              Text(
                "Fluent use",
                style: TextStyle(color: Color.fromRGBO(160, 160, 160, 1)),
              ),
              Text(
                "10 words",
                style: TextStyle(color: Color.fromRGBO(100, 100, 100, 1)),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
