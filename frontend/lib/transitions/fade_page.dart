import 'package:flutter/material.dart';

class FadeTransitionRoute extends PageRouteBuilder {
  final Widget child;

  FadeTransitionRoute({
    required this.child,
  }) : super(
          transitionDuration: const Duration(milliseconds: 150),
          reverseTransitionDuration: const Duration(milliseconds: 150),
          pageBuilder: (context, animation, secondaryAnimation) => child,
        );

  @override
  Widget buildTransitions(BuildContext context, Animation<double> animation,
          Animation<double> secondaryAnimation, Widget child) =>
      FadeTransition(
        opacity: animation,
        child: child,
      );
}
