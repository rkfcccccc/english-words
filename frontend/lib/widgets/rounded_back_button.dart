import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class RoundedBackButton extends StatelessWidget {
  const RoundedBackButton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () => Navigator.of(context).pop(),
      style: ElevatedButton.styleFrom(
        padding: EdgeInsets.all(3.5.w),
        primary: Colors.transparent,
        onPrimary: const Color.fromRGBO(54, 54, 54, 1),
        minimumSize: Size.zero,
        side: const BorderSide(
          width: 2,
          color: Color.fromRGBO(33, 33, 33, 1),
        ),
      ),
      child: Icon(
        Icons.arrow_back_ios_new_rounded,
        size: 13.sp,
        color: const Color.fromRGBO(212, 212, 212, 1),
      ),
    );
  }
}
