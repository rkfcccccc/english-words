import 'package:english_words/screens/signup.dart';
import 'package:english_words/transitions/fade_page.dart';
import 'package:english_words/transitions/no_animation.dart';
import 'package:english_words/utils/errors.dart';
import 'package:english_words/widgets/gradient_button.dart';
import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:english_words/utils/api/user/user.dart' as user_api;
import 'package:english_words/utils/api/user/auth.dart' as auth_api;
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'home.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  String email = "", password = "";

  void onContinue() async {
    try {
      final navigator = Navigator.of(context);
      final data = await user_api.login(email, password);
      await auth_api.storeCredentials(data.jwt, data.refresh);

      navigator.pushReplacement(
        FadeTransitionRoute(child: const HomeScreen()),
      );
    } on AppError catch (err) {
      final snackBar = SnackBar(
        content: Text(err.message ?? "Unknown error"),
      );

      ScaffoldMessenger.of(context).showSnackBar(snackBar);
    }
  }

  bool get canContinue => email != "" && password != "";

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
                validator: (x) => x != "",
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
                validator: (x) => x != "",
                onChanged: (text) => setState(() {
                  password = text;
                }),
                onSubmitted: (_) => canContinue ? onContinue() : null,
                obscureText: true,
              ),
              SizedBox(height: indent),
              GradientButton(
                onPressed: canContinue ? onContinue : null,
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
                      NoAnimationRoute(child: const SignupScreen()),
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
