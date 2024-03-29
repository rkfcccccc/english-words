import 'package:english_words/screens/home.dart';
import 'package:english_words/screens/signup.dart';
import 'package:english_words/transitions/fade_page.dart';
import 'package:english_words/utils/api/user/auth.dart';
import 'package:flutter/material.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({Key? key}) : super(key: key);

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  void routeSomewhere() async {
    final navigator = Navigator.of(context);
    final token = await getAccessToken();

    navigator.pushReplacement(
      FadeTransitionRoute(
        child: token != null ? const HomeScreen() : const SignupScreen(),
      ),
    );
  }

  @override
  void initState() {
    routeSomewhere();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return const Scaffold();
  }
}
