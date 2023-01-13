import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'floating_container.dart';

class FloatingWrapper extends StatefulWidget {
  final Widget child;
  final Widget? actions;

  const FloatingWrapper({
    Key? key,
    required this.child,
    this.actions,
  }) : super(key: key);

  @override
  State<FloatingWrapper> createState() => _FloatingWrapperState();
}

class _FloatingWrapperState extends State<FloatingWrapper> {
  late ScrollController _controller;
  bool floating = false;

  void updateFloating() {
    assert(_controller.position.hasContentDimensions);

    if (_controller.position.maxScrollExtent != 0) {
      if (!floating) {
        setState(() {
          floating = true;
        });
      }
    } else {
      if (floating) {
        setState(() {
          floating = false;
        });
      }
    }
  }

  @override
  void initState() {
    _controller = ScrollController();
    super.initState();

    WidgetsBinding.instance.addPostFrameCallback((_) => updateFloating());
  }

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: SingleChildScrollView(
        clipBehavior: Clip.none,
        controller: _controller,
        physics: const AlwaysScrollableScrollPhysics(),
        child: Padding(
          padding: EdgeInsets.all(6.w),
          child: Column(
            children: [
              FloatingContainer(
                child: Column(
                  children: [
                    widget.child,
                    if (widget.actions != null && !floating) ...[
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 4.w),
                        child: const Divider(height: 0),
                      ),
                      widget.actions!,
                    ]
                  ],
                ),
              ),
              if (widget.actions != null && floating) ...[
                SizedBox(height: 6.w),
                FloatingContainer(child: widget.actions!),
              ]
            ],
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
    return Padding(
      padding: EdgeInsets.all(4.w),
      child: Column(
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
      ),
    );
  }
}
