import 'package:english_words/styles/gradients.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class GradientButton extends StatelessWidget {
  final VoidCallback? onPressed;
  final String text;
  final bool loading;

  const GradientButton(
      {Key? key,
      required this.onPressed,
      required this.text,
      this.loading = false})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final textWidget = Text(
      text,
      style: TextStyle(
        fontSize: 16,
        color: onPressed != null ? null : const Color.fromRGBO(40, 40, 40, 1),
        fontWeight: FontWeight.w600,
      ),
    );

    final loadingWidget = SizedBox.square(
      dimension: 3.w,
      child: const CircularProgressIndicator(
        strokeWidth: 2,
        color: Colors.white,
      ),
    );

    return Container(
      width: double.infinity,
      height: 12.w,
      decoration: BoxDecoration(
        gradient: onPressed != null ? Gradients.purple2pink : null,
        color: onPressed != null ? null : const Color.fromRGBO(23, 23, 23, 1),
        borderRadius: BorderRadius.circular(13),
      ),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(
          primary: Colors.transparent,
          onSurface: Colors.transparent,
        ),
        onPressed: onPressed,
        child: loading ? loadingWidget : textWidget,
      ),
    );
  }
}
