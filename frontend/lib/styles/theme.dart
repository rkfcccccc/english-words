import 'package:flutter/material.dart';

import 'gradients.dart';

var appTheme = ThemeData.dark().copyWith(
  textSelectionTheme: TextSelectionThemeData(
      selectionColor: Gradients.purple2pink.colors.last.withOpacity(0.2)),
  scaffoldBackgroundColor: Color.fromRGBO(17, 17, 17, 1),
  elevatedButtonTheme: ElevatedButtonThemeData(
    style: ButtonStyle(
      elevation: MaterialStateProperty.all(0),
      shape: MaterialStateProperty.all(
        RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10),
        ),
      ),
      padding: MaterialStateProperty.all(
        const EdgeInsets.symmetric(vertical: 12.0, horizontal: 36),
      ),
    ),
  ),
  textButtonTheme: TextButtonThemeData(
    style: ButtonStyle(
      shape: MaterialStateProperty.all(
        RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(7),
        ),
      ),
    ),
  ),
);
