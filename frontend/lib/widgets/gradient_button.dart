import 'package:english_words/styles/gradients.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class GradientButton extends StatelessWidget {
  const GradientButton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 12.w,
      decoration: BoxDecoration(
        gradient: Gradients.purple2pink,
        borderRadius: BorderRadius.circular(13),
      ),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(primary: Colors.transparent),
        onPressed: () {},
        child: const Text(
          "Continue",
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w600,
          ),
        ),
      ),
    );
  }
}
