import 'package:english_words/screens/login.dart';
import 'package:english_words/screens/verification.dart';
import 'package:english_words/transitions/no_animation.dart';
import 'package:english_words/widgets/gradient_button.dart';
import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

final emailRegexp = RegExp(
    "^[a-zA-Z0-9.!#\$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\$");

class SignupScreen extends StatefulWidget {
  const SignupScreen({Key? key}) : super(key: key);

  @override
  State<SignupScreen> createState() => _SignupScreenState();
}

class _SignupScreenState extends State<SignupScreen> {
  String email = "", password = "";

  bool emailValidator(String text) {
    return text.length <= 64 && emailRegexp.matchAsPrefix(text) != null;
  }

  bool passwordValidator(String text) {
    return text.length <= 72 && text.isNotEmpty;
  }

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
    final canContinue = emailValidator(email) && passwordValidator(password);

    return Scaffold(
      body: Padding(
        padding: EdgeInsets.all(indent),
        child: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Sign up",
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
                validator: emailValidator,
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
                onSubmitted: (_) => canContinue ? onContinue() : null,
                validator: passwordValidator,
                obscureText: true,
              ),
              SizedBox(height: indent),
              GradientButton(
                // loading: true,
                onPressed: canContinue ? onContinue : null,
                text: "Continue",
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "Already have an account?",
                    style: TextStyle(
                        fontSize: 10.sp,
                        color: const Color.fromRGBO(142, 142, 142, 1)),
                  ),
                  CupertinoButton(
                    onPressed: () => Navigator.of(context).pushReplacement(
                      NoAnimationRoute(child: const LoginScreen()),
                    ),
                    child: Text(
                      "Log in",
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
