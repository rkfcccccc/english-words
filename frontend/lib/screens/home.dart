import 'dart:ui';

import 'package:english_words/widgets/section.dart';
import 'package:english_words/widgets/section_favorite_movies.dart';
import 'package:english_words/widgets/section_vocabulary.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.all(6.w),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "Hello!",
                  style: TextStyle(
                    fontSize: 23.sp,
                    fontWeight: FontWeight.w700,
                  ),
                ),
                SizedBox(height: 4.w),
                const VocabularySection(),
                SizedBox(height: 4.w),
                const FavoriteMoviesSection(),
                SizedBox(height: 4.w),
                Section(
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: const [
                      Text(
                        "Start lesson",
                        style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                      Icon(Icons.chevron_right_rounded)
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
