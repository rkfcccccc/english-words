import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:english_words/widgets/rounded_back_button.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class SearchScreen extends StatelessWidget {
  const SearchScreen({Key? key}) : super(key: key);

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
                    "Search",
                    style: TextStyle(
                      fontSize: 23.sp,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ],
              ),
              SizedBox(height: 4.w),
              GradientTextField(
                padding: EdgeInsets.symmetric(horizontal: 4.w, vertical: 3.w),
                autofocus: true,
                placeholder: "What are you looking for?",
              ),
              SizedBox(height: 2.w),
              Row(
                children: [
                  const _SearchTab(text: "words"),
                  SizedBox(width: 2.w),
                  const _SearchTab(text: "movies"),
                ],
              ),
              SizedBox(height: 2.w),
              const _Item(child: _WordItem()),
              SizedBox(height: 2.w),
              const _Item(child: _MovieItem()),
            ],
          ),
        ),
      ),
    );
  }
}

class _MovieItem extends StatelessWidget {
  const _MovieItem({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(7),
            image: const DecorationImage(
              image: NetworkImage(
                  "https://i.jeded.com/i/harry-potter-and-the-goblet-of-fire.200-11171.jpg"),
              fit: BoxFit.fill,
            ),
          ),
          width: 20.w,
          height: 20.w / 2 * 3,
          clipBehavior: Clip.antiAlias,
        ),
        SizedBox(width: 4.w),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "Harry Potter",
                  style: TextStyle(
                    fontSize: 12.sp,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                Text(
                  "2022",
                  style: TextStyle(
                    fontSize: 10.sp,
                    color: const Color.fromRGBO(160, 160, 160, 1),
                  ),
                ),
              ],
            ),
            // SizedBox(height: 2.w),
            Container(
              padding: EdgeInsets.all(1.w),
              decoration: BoxDecoration(
                color: Colors.purple,
                borderRadius: BorderRadius.circular(7),
              ),
              child: const Text("78%"),
            )
          ],
        )
      ],
    );
  }
}

class _Item extends StatelessWidget {
  final Widget child;
  const _Item({
    Key? key,
    required this.child,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: EdgeInsets.all(4.w),
      decoration: BoxDecoration(
        color: const Color.fromRGBO(23, 23, 23, 1),
        borderRadius: BorderRadius.circular(13),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [child, const Icon(Icons.chevron_right_rounded)],
      ),
    );
  }
}

class _WordItem extends StatelessWidget {
  const _WordItem({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "definitions of",
          style: TextStyle(
            color: const Color.fromRGBO(140, 140, 140, 1),
            fontSize: 10.sp,
          ),
        ),
        Text(
          "dungeon",
          style: TextStyle(
            fontWeight: FontWeight.w500,
            fontSize: 12.sp,
          ),
        ),
      ],
    );
  }
}

class _SearchTab extends StatelessWidget {
  final String text;

  const _SearchTab({
    Key? key,
    required this.text,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 20.w,
      height: 9.w,
      child: ElevatedButton(
        onPressed: () {},
        style: ElevatedButton.styleFrom(
          minimumSize: Size.zero,
          elevation: 0,
          primary: const Color.fromRGBO(23, 23, 23, 1),
          onPrimary: const Color.fromRGBO(54, 54, 54, 1),
          onSurface: const Color.fromRGBO(122, 122, 122, 1),
          padding: EdgeInsets.zero,
        ),
        child: Text(
          text,
          style: TextStyle(
            fontSize: 11.sp,
            color: const Color.fromRGBO(222, 222, 222, 1),
          ),
        ),
      ),
    );
  }
}
