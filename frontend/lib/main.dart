import 'package:english_words/styles/theme.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'screens/signup.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Sizer(
      builder: (context, orientation, deviceType) => MaterialApp(
        themeMode: ThemeMode.light,
        theme: appTheme,
        home: const SignupScreen(),
      ),
    );
  }
}
