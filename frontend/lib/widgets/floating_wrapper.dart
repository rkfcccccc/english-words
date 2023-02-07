import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'floating_container.dart';

class FloatingWrapper extends StatelessWidget {
  final List<Widget> children;

  const FloatingWrapper({
    Key? key,
    required this.children,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: SingleChildScrollView(
        clipBehavior: Clip.none,
        physics: const AlwaysScrollableScrollPhysics(),
        child: Padding(
          padding: EdgeInsets.all(6.w),
          child: Column(
            children: children
                .asMap()
                .entries
                .map(
                  (e) => Padding(
                    padding: e.key == children.length - 1
                        ? EdgeInsets.zero
                        : EdgeInsets.only(bottom: 6.w),
                    child: FloatingContainer(
                      child: e.value,
                    ),
                  ),
                )
                .toList(),
          ),
        ),
      ),
    );
  }
}

class FloatingActions extends StatelessWidget {
  final VoidCallback? onPrimary, onSecondary;
  final String? primaryText, secondaryText;

  const FloatingActions({
    Key? key,
    this.onPrimary,
    this.onSecondary,
    this.primaryText,
    this.secondaryText,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        if (onPrimary != null)
          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: onPrimary,
              child: Text(primaryText!, style: TextStyle(fontSize: 12.sp)),
            ),
          ),
        if (onPrimary != null && onSecondary != null) SizedBox(height: 1.w),
        if (onSecondary != null)
          SizedBox(
            width: double.infinity,
            height: 10.w,
            child: TextButton(
              onPressed: onSecondary,
              child: Text(
                secondaryText!,
                style: TextStyle(fontSize: 11.sp),
              ),
            ),
          ),
      ],
    );
  }
}
