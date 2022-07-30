import 'package:english_words/widgets/gradient_button.dart';
import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class SignupScreen extends StatelessWidget {
  const SignupScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final indent = 6.w;

    return Scaffold(
      body: Padding(
        padding: EdgeInsets.all(6.w),
        child: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                "Sign up",
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.w700,
                ),
              ),
              SizedBox(height: indent),
              const GradientTextField(
                label: "Email",
                placeholder: "example@gmail.com",
              ),
              SizedBox(height: indent),
              const GradientTextField(
                label: "Password",
                placeholder: "very_strong_password",
                obscureText: true,
              ),
              SizedBox(height: indent),
              const GradientButton(),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text(
                    "Already have an account?",
                    style: TextStyle(color: Color.fromRGBO(142, 142, 142, 1)),
                  ),
                  CupertinoButton(
                    onPressed: () {},
                    child: const Text(
                      "Log in",
                      style: TextStyle(
                          fontSize: 14,
                          color: Color.fromRGBO(222, 222, 222, 1)),
                    ),
                  ),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }
}
