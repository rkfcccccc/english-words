import 'package:english_words/styles/gradients.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class _Painter extends CustomPainter {
  final Gradient gradient;
  final Radius? radius;
  final double strokeWidth;
  final double opacity;

  _Painter(this.gradient, this.radius, this.strokeWidth, this.opacity);

  @override
  void paint(Canvas canvas, Size size) {
    final rect = Rect.fromLTWH(strokeWidth / 2, strokeWidth / 2,
        size.width - strokeWidth, size.height - strokeWidth);
    final rRect = RRect.fromRectAndRadius(rect, radius ?? Radius.zero);

    final paint = Paint()
      ..style = PaintingStyle.stroke
      ..strokeWidth = strokeWidth
      ..color = Colors.white.withOpacity(opacity)
      ..shader = gradient.createShader(rect);

    canvas.drawRRect(rRect, paint);
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) => oldDelegate != this;
}

class GradientTextField extends StatefulWidget {
  final String label;
  final String placeholder;
  final bool obscureText;

  const GradientTextField(
      {Key? key,
      required this.label,
      required this.placeholder,
      this.obscureText = false})
      : super(key: key);

  @override
  State<GradientTextField> createState() => _GradientTextFieldState();
}

class _GradientTextFieldState extends State<GradientTextField>
    with SingleTickerProviderStateMixin {
  late Animation animation;
  late AnimationController _controller;
  late FocusNode _focusNode;

  @override
  void initState() {
    _controller = AnimationController(
        duration: const Duration(milliseconds: 150), vsync: this);

    animation = Tween<double>(begin: 0, end: 1).animate(_controller)
      ..addListener(() {
        setState(() {});
      });

    _focusNode = FocusNode()
      ..addListener(() {
        setState(() {
          if (_focusNode.hasFocus) {
            _controller.forward();
          } else {
            _controller.reverse();
          }
        });
      });

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => _focusNode.requestFocus(),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.only(left: 13),
            child: Text(widget.label,
                style: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                    color: Color.fromRGBO(225, 225, 225, 1))),
          ),
          SizedBox(height: 2.w),
          Container(
            decoration: BoxDecoration(
                color: const Color.fromRGBO(23, 23, 23, 1),
                borderRadius: BorderRadius.circular(13),
                border: Border.all(color: const Color.fromRGBO(30, 30, 30, 1))),
            child: CustomPaint(
              painter: _Painter(Gradients.purple2pink,
                  const Radius.circular(12), 2, animation.value),
              child: Padding(
                padding: EdgeInsets.all(4.w),
                child: SizedBox(
                  child: TextField(
                    focusNode: _focusNode,
                    autocorrect: false,
                    obscureText: widget.obscureText,
                    cursorColor: Gradients.purple2pink.colors.first,
                    decoration: InputDecoration.collapsed(
                      hintText: widget.placeholder,
                    ),
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
