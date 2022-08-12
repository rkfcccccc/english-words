import 'package:flutter/material.dart';
import 'gradients.dart';

var appTheme = ThemeData.dark().copyWith(
  textSelectionTheme: TextSelectionThemeData(
      selectionColor: Gradients.purple2pink.colors.last.withOpacity(0.2)),
  scaffoldBackgroundColor: const Color.fromRGBO(17, 17, 17, 1),
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
  snackBarTheme: SnackBarThemeData(
    elevation: 0,
    behavior: SnackBarBehavior.floating,
    backgroundColor: const Color.fromRGBO(33, 33, 33, 1),
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(13)),
    contentTextStyle: const TextStyle(
      color: Color.fromRGBO(205, 205, 205, 1),
      fontSize: 15,
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
