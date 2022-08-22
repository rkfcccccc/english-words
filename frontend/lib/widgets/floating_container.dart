import 'package:flutter/material.dart';

class FloatingContainer extends StatelessWidget {
  final Widget? child;
  const FloatingContainer({Key? key, this.child}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      clipBehavior: Clip.antiAlias,
      decoration: BoxDecoration(
        color: const Color.fromRGBO(32, 32, 32, 1),
        borderRadius: BorderRadius.circular(13),
        boxShadow: [
          BoxShadow(
            spreadRadius: 1,
            blurRadius: 12,
            color: const Color.fromRGBO(32, 32, 32, 1).withOpacity(0.25),
          )
        ],
      ),
      child: child,
    );
  }
}
