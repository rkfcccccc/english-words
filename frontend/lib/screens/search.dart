import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:english_words/widgets/rounded_back_button.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class SearchScreen extends StatelessWidget {
  const SearchScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: EdgeInsets.all(6.w),
        child: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  const RoundedBackButton(),
                  SizedBox(width: 6.w),
                  Text(
                    "Search",
                    style: TextStyle(
                      fontSize: 23.sp,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ],
              ),
              SizedBox(height: 4.w),
              GradientTextField(
                padding: EdgeInsets.symmetric(horizontal: 4.w, vertical: 3.w),
                autofocus: true,
                placeholder: "What are you looking for?",
              ),
              SizedBox(height: 4.w),
              Container(
                width: double.infinity,
                padding: EdgeInsets.all(4.w),
                decoration: BoxDecoration(
                  color: Color.fromRGBO(23, 23, 23, 1),
                  borderRadius: BorderRadius.circular(13),
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
                            color: Color.fromRGBO(140, 140, 140, 1),
                            fontSize: 10.sp,
                          ),
                        ),
                        Text(
                          "dungeon",
                          style: TextStyle(
                            fontWeight: FontWeight.w500,
                            fontSize: 12.sp,
                          ),
                        ),
                      ],
                    ),
                    const Icon(Icons.chevron_right_rounded)
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
