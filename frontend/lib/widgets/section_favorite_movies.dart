import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

import 'section.dart';

class FavoriteMoviesSection extends StatelessWidget {
  const FavoriteMoviesSection({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Section(
      header: Text(
        'Favorite movies',
        style: TextStyle(
          fontSize: 11.sp,
          fontWeight: FontWeight.w600,
        ),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(13),
              image: const DecorationImage(
                image: NetworkImage(
                    "https://i.jeded.com/i/harry-potter-and-the-goblet-of-fire.200-11171.jpg"),
                fit: BoxFit.fill,
              ),
            ),
            width: 25.w,
            height: 25.w / 2 * 3,
            clipBehavior: Clip.antiAlias,
          ),
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(13),
              image: const DecorationImage(
                image: NetworkImage(
                    "https://i.jeded.com/i/star-wars-episode-i--the-phantom-menace.200-13951.jpg"),
                fit: BoxFit.fill,
              ),
            ),
            width: 25.w,
            height: 25.w / 2 * 3,
            clipBehavior: Clip.antiAlias,
          ),
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(13),
              image: const DecorationImage(
                image: NetworkImage(
                    "https://i.jeded.com/i/doctor-strange-2016.200-179773.jpg"),
                fit: BoxFit.fill,
              ),
            ),
            width: 25.w,
            height: 25.w / 2 * 3,
            clipBehavior: Clip.antiAlias,
          )
        ],
      ),
    );
  }
}
