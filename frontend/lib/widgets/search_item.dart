import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class SearchItem extends StatelessWidget {
  final Widget child;
  final VoidCallback? onPressed;

  const SearchItem({
    Key? key,
    required this.child,
    this.onPressed,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed,
      style: ElevatedButton.styleFrom(
        alignment: Alignment.centerLeft,
        primary: const Color.fromRGBO(23, 23, 23, 1),
        onPrimary: const Color.fromRGBO(54, 54, 54, 1),
        padding: EdgeInsets.all(4.w),
        textStyle: const TextStyle(),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(13)),
      ),
      child: child,
    );
  }
}
