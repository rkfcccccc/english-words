import 'package:flutter/material.dart';

class NoAnimationRoute extends PageRouteBuilder {
  final Widget child;

  NoAnimationRoute({
    required this.child,
  }) : super(
          transitionDuration: const Duration(milliseconds: 0),
          pageBuilder: (context, animation, secondaryAnimation) => child,
        );
}
