import 'package:english_words/screens/signup.dart';
import 'package:english_words/screens/verification.dart';
import 'package:english_words/widgets/gradient_button.dart';
import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  String email = "", password = "";

  void onContinue() {
    Navigator.of(context).push(
      MaterialPageRoute(
        builder: (ctx) => VerificationScreen(
          email: email,
          onSubmitted: print,
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    final indent = 6.w;

    return Scaffold(
      body: Padding(
        padding: EdgeInsets.all(indent),
        child: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Log in",
                style: TextStyle(
                  fontSize: 23.sp,
                  fontWeight: FontWeight.w700,
                ),
              ),
              SizedBox(height: indent),
              GradientTextField(
                label: "Email",
                placeholder: "example@gmail.com",
                onChanged: (text) => setState(() {
                  email = text;
                }),
                keyboardType: TextInputType.emailAddress,
                textInputAction: TextInputAction.next,
                autofocus: true,
              ),
              SizedBox(height: indent),
              GradientTextField(
                label: "Password",
                placeholder: "very_strong_password",
                onChanged: (text) => setState(() {
                  password = text;
                }),
                onSubmitted: print,
                obscureText: true,
              ),
              SizedBox(height: indent),
              GradientButton(
                onPressed: () {},
                text: "Log in",
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "Don't have an account?",
                    style: TextStyle(
                        fontSize: 10.sp,
                        color: const Color.fromRGBO(142, 142, 142, 1)),
                  ),
                  CupertinoButton(
                    onPressed: () => Navigator.of(context).pushReplacement(
                      MaterialPageRoute(
                        builder: (ctx) => const SignupScreen(),
                        fullscreenDialog: true,
                      ),
                    ),
                    child: Text(
                      "Sign up",
                      style: TextStyle(
                          fontSize: 10.sp,
                          color: const Color.fromRGBO(222, 222, 222, 1)),
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
