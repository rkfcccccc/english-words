import 'dart:async';

import 'package:english_words/utils/errors.dart';
import 'package:english_words/widgets/rounded_back_button.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:sizer/sizer.dart';

const resendSeconds = 60;

class VerificationScreen extends StatefulWidget {
  final String email;
  final String requestId;
  final Future<void> Function(String requestId, int code) onSubmitted;
  final Future<String> Function()? onResended;

  const VerificationScreen({
    Key? key,
    required this.email,
    required this.requestId,
    required this.onSubmitted,
    this.onResended,
  }) : super(key: key);

  @override
  State<VerificationScreen> createState() => _VerificationScreenState();
}

class _VerificationScreenState extends State<VerificationScreen> {
  late String requestId;
  late Timer timer;

  late int startTime;
  int secondsLeft = resendSeconds;

  @override
  void initState() {
    requestId = widget.requestId;

    startTime = DateTime.now().millisecondsSinceEpoch;
    timer = Timer.periodic(const Duration(seconds: 1), (_) {
      final secondsPassed =
          (DateTime.now().millisecondsSinceEpoch - startTime) ~/ 1000;

      setState(() {
        secondsLeft = resendSeconds - secondsPassed;
      });
    });

    super.initState();
  }

  Future<bool> onSubmit(int code) async {
    try {
      await widget.onSubmitted(requestId, code);
      return true;
    } on AppError catch (err) {
      final snackBar = SnackBar(
        behavior: SnackBarBehavior.floating,
        content: Text(err.message ?? "Unknown error"),
      );

      ScaffoldMessenger.of(context).showSnackBar(snackBar);
    }

    return false;
  }

  void onResend() async {
    setState(() {
      startTime = DateTime.now().millisecondsSinceEpoch;
      secondsLeft = resendSeconds;
    });

    requestId = await widget.onResended!();
  }

  @override
  void dispose() {
    timer.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: EdgeInsets.all(6.w),
        child: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  const RoundedBackButton(),
                  SizedBox(width: 6.w),
                  Text(
                    "Verification",
                    style: TextStyle(
                      fontSize: 23.sp,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ],
              ),
              SizedBox(height: 6.w),
              RichText(
                text: TextSpan(
                  text: "We've sent you a 4-digit code to ",
                  style: TextStyle(
                    color: const Color.fromRGBO(242, 242, 242, 1),
                    fontSize: 10.sp,
                  ),
                  children: <TextSpan>[
                    TextSpan(
                      text: widget.email,
                      style: const TextStyle(
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                    const TextSpan(
                      text:
                          '. Enter it to confirm that you are the owner of this email.',
                    ),
                  ],
                ),
              ),
              SizedBox(height: 6.w),
              _OtpCodeInput(onSubmit: onSubmit),
              if (widget.onResended != null)
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    const Text(
                      "Did not get the code?",
                      style: TextStyle(color: Color.fromRGBO(142, 142, 142, 1)),
                    ),
                    CupertinoButton(
                      onPressed: secondsLeft <= 0 ? onResend : null,
                      child: Text(
                        "Resend${secondsLeft <= 0 ? "" : " in $secondsLeft seconds"}",
                        style: const TextStyle(fontSize: 14),
                      ),
                    ),
                  ],
                ),
            ],
          ),
        ),
      ),
    );
  }
}

class _OtpCodeInput extends StatefulWidget {
  final Future<bool> Function(int)? onSubmit;

  const _OtpCodeInput({
    Key? key,
    this.onSubmit,
  }) : super(key: key);

  @override
  State<_OtpCodeInput> createState() => _OtpCodeInputState();
}

class _OtpCodeInputState extends State<_OtpCodeInput> {
  late TextEditingController _controller;
  late FocusNode _focusNode;

  String code = "";

  @override
  void initState() {
    _controller = TextEditingController();
    _focusNode = FocusNode();

    super.initState();
  }

  void onChanged(String text) async {
    setState(() {
      code = text;
    });

    if (text.length == 4 && widget.onSubmit != null) {
      _focusNode.unfocus();
      final ok = await widget.onSubmit!(int.parse(text));

      if (!ok) {
        _focusNode.requestFocus();
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => _focusNode.requestFocus(),
      child: Stack(
        children: [
          Visibility(
            visible: false,
            maintainSize: true,
            maintainState: true,
            maintainAnimation: true,
            child: TextField(
              autofocus: true,
              controller: _controller,
              focusNode: _focusNode,
              keyboardType: TextInputType.number,
              keyboardAppearance: Brightness.dark,
              inputFormatters: [
                LengthLimitingTextInputFormatter(4),
                FilteringTextInputFormatter.digitsOnly
              ],
              onChanged: onChanged,
            ),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              _OtpDigit(digit: code.isEmpty ? "" : code[0]),
              _OtpDigit(digit: code.length <= 1 ? "" : code[1]),
              _OtpDigit(digit: code.length <= 2 ? "" : code[2]),
              _OtpDigit(digit: code.length <= 3 ? "" : code[3]),
            ],
          ),
        ],
      ),
    );
  }
}

class _OtpDigit extends StatelessWidget {
  final String digit;

  const _OtpDigit({Key? key, required this.digit}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 15.w,
      height: 15.w,
      decoration: BoxDecoration(
        color: const Color.fromRGBO(23, 23, 23, 1),
        borderRadius: BorderRadius.circular(13),
      ),
      child: Center(
        child: Text(digit, style: Theme.of(context).textTheme.headline6),
      ),
    );
  }
}
