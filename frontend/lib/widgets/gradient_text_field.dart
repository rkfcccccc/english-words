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
  final String? label;
  final String placeholder;
  final void Function(String)? onSubmitted;
  final void Function(String)? onChanged;
  final bool Function(String)? validator;
  final VoidCallback? onFocused;
  final TextInputType? keyboardType;
  final TextInputAction? textInputAction;
  final bool obscureText;
  final bool autofocus;
  final EdgeInsets? padding;

  const GradientTextField({
    Key? key,
    required this.placeholder,
    this.label,
    this.onFocused,
    this.padding,
    this.obscureText = false,
    this.autofocus = false,
    this.textInputAction,
    this.keyboardType,
    this.onSubmitted,
    this.onChanged,
    this.validator,
  }) : super(key: key);

  @override
  State<GradientTextField> createState() => _GradientTextFieldState();
}

class _GradientTextFieldState extends State<GradientTextField>
    with SingleTickerProviderStateMixin {
  late TextEditingController _editingController;
  late AnimationController _controller;
  late Animation animation;
  late FocusNode _focusNode;

  bool _touched = false;
  bool _passesValidator = true;

  @override
  void initState() {
    _editingController = TextEditingController()
      ..addListener(() {
        if (widget.validator != null) {
          _passesValidator = widget.validator!(_editingController.text);
        }
      });

    _controller = AnimationController(
        duration: const Duration(milliseconds: 150), vsync: this);

    animation = Tween<double>(begin: 0, end: 1).animate(_controller)
      ..addListener(() {
        setState(() {});
      });

    if (widget.autofocus) {
      _controller.forward(from: 1);
    }

    _focusNode = FocusNode()
      ..addListener(() {
        setState(() {
          if (_focusNode.hasFocus) {
            if (widget.onFocused != null) widget.onFocused!();

            _controller.forward();
            _touched = true;
          } else {
            _controller.reverse();
          }
        });
      });

    super.initState();
  }

  @override
  void dispose() {
    _focusNode.dispose();
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    var showValidator =
        _touched && !_focusNode.hasFocus && widget.validator != null;

    return GestureDetector(
      onTap: () => _focusNode.requestFocus(),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          if (widget.label != null)
            Padding(
              padding: EdgeInsets.only(left: 13, bottom: 2.w),
              child: Text(
                widget.label!,
                style: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                    color: Color.fromRGBO(225, 225, 225, 1)),
              ),
            ),
          Container(
            decoration: BoxDecoration(
              color: const Color.fromRGBO(23, 23, 23, 1),
              borderRadius: BorderRadius.circular(13),
              border: Border.all(
                color: const Color.fromRGBO(33, 33, 33, 1),
                width: 2,
              ),
            ),
            child: CustomPaint(
              painter: _Painter(Gradients.purple2pink,
                  const Radius.circular(12), 2, animation.value),
              child: SizedBox(
                child: Padding(
                  padding: widget.padding ?? EdgeInsets.all(4.w),
                  child: Row(
                    children: [
                      Flexible(
                        child: TextField(
                          autocorrect: false,
                          keyboardType: widget.keyboardType,
                          autofocus: widget.autofocus,
                          onChanged: widget.onChanged,
                          obscureText: widget.obscureText,
                          textInputAction: widget.textInputAction,
                          onSubmitted: widget.onSubmitted,
                          controller: _editingController,
                          focusNode: _focusNode,
                          cursorColor: Gradients.purple2pink.colors.first,
                          decoration: InputDecoration.collapsed(
                            hintText: widget.placeholder,
                          ),
                        ),
                      ),
                      AnimatedOpacity(
                        opacity: showValidator ? 1 : 0,
                        duration: const Duration(milliseconds: 150),
                        child: RoundedValidatorStatus(ok: _passesValidator),
                      )
                    ],
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

class RoundedValidatorStatus extends StatelessWidget {
  final bool ok;

  const RoundedValidatorStatus({
    Key? key,
    required this.ok,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        SizedBox(width: 2.w),
        Container(
          padding: EdgeInsets.all(1.w),
          decoration: BoxDecoration(
            color: const Color.fromRGBO(33, 33, 33, 1),
            borderRadius: BorderRadius.circular(6),
          ),
          child: Icon(
            ok ? Icons.check_rounded : Icons.close_rounded,
            size: 10.sp,
          ),
        )
      ],
    );
  }
}
