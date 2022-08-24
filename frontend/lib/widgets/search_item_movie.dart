import 'package:english_words/utils/api/movies/models.dart';
import 'package:english_words/widgets/search_item.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class MovieSearchItem extends StatelessWidget {
  final Movie movie;

  const MovieSearchItem({
    Key? key,
    required this.movie,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SearchItem(
      onPressed: () {},
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(7),
              image: DecorationImage(
                image: NetworkImage(movie.posterUrl),
                fit: BoxFit.fill,
              ),
            ),
            width: 20.w,
            height: 20.w / 2 * 3,
            clipBehavior: Clip.antiAlias,
          ),
          SizedBox(width: 4.w),
          Flexible(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      movie.title,
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 12.sp,
                        fontWeight: FontWeight.w500,
                      ),
                    ),
                    Text(
                      movie.year.toString(),
                      style: TextStyle(
                        fontSize: 10.sp,
                        color: const Color.fromRGBO(160, 160, 160, 1),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}
