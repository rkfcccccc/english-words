import 'package:english_words/screens/home.dart';
import 'package:english_words/screens/login.dart';
import 'package:english_words/screens/verification.dart';
import 'package:english_words/transitions/fade_page.dart';
import 'package:english_words/transitions/no_animation.dart';
import 'package:english_words/utils/api/user/models.dart';
import 'package:english_words/utils/api/user/user.dart' as user_api;
import 'package:english_words/utils/errors.dart';
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
  bool loading = false;

  bool emailValidator(String text) {
    return text.length <= 64 && emailRegexp.matchAsPrefix(text) != null;
  }

  bool passwordValidator(String text) {
    return text.length >= 6 && text.length <= 72 && text.isNotEmpty;
  }

  Future<String> proceedVerification() async {
    final data = await user_api.signup(email, password);
    return data.requestId!;
  }

  Future<void> onVerificationSubmitted(
    String requestId,
    int code,
  ) async {
    final navigator = Navigator.of(context);
    final data = await user_api.signup(
      email,
      password,
      verification: VerificationData(
        requestId: requestId,
        code: code,
      ),
    );

    await user_api.storeCredentials(data.jwt!, data.refresh!);

    navigator.pushReplacement(
      FadeTransitionRoute(child: const HomeScreen()),
    );
  }

  void onContinue() async {
    setState(() {
      loading = true;
    });

    try {
      final navigator = Navigator.of(context);
      final requestId = await proceedVerification();

      navigator.push(
        MaterialPageRoute(
          builder: (ctx) => VerificationScreen(
            email: email,
            requestId: requestId,
            onResended: proceedVerification,
            onSubmitted: onVerificationSubmitted,
          ),
        ),
      );
    } on AppError catch (err) {
      final snackBar = SnackBar(
        behavior: SnackBarBehavior.floating,
        content: Text(err.message ?? "Unknown error"),
      );

      ScaffoldMessenger.of(context).showSnackBar(snackBar);
    }

    await Future.delayed(const Duration(milliseconds: 250));

    setState(() {
      loading = false;
    });
  }

  @override
  void initState() {
    super.initState();
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
                loading: loading,
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
