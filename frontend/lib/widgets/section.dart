import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class Section extends StatelessWidget {
  final Widget? header;
  final Widget? child;
  const Section({Key? key, this.header, this.child}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      decoration: BoxDecoration(
        color: const Color.fromRGBO(20, 20, 20, 1),
        borderRadius: BorderRadius.circular(13),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          if (header != null)
            Container(
              width: double.infinity,
              padding: EdgeInsets.all(3.w),
              decoration: BoxDecoration(
                color: const Color.fromRGBO(23, 23, 23, 1),
                borderRadius: BorderRadius.circular(13),
              ),
              child: header,
            ),
          if (child != null)
            Padding(
              padding: EdgeInsets.all(3.w),
              child: child,
            )
        ],
      ),
    );
  }
}
