import 'package:english_words/utils/api/dictionary/models.dart';
import 'package:english_words/utils/api/movies/models.dart';
import 'package:english_words/utils/api/movies/movies.dart' as movies_api;
import 'package:english_words/utils/api/dictionary/dictionary.dart'
    as dictionary_api;
import 'package:english_words/widgets/floating_container.dart';
import 'package:english_words/widgets/floating_wrapper.dart';
import 'package:english_words/widgets/gradient_text_field.dart';
import 'package:english_words/widgets/rounded_back_button.dart';
import 'package:english_words/widgets/search_item_movie.dart';
import 'package:english_words/widgets/search_item_word.dart';
import 'package:english_words/widgets/word_entry_view.dart';
import 'package:flutter/material.dart';
import 'package:sizer/sizer.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({Key? key}) : super(key: key);

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  int _selectedTab = 1;
  String query = "";

  void selectTab(int index) {
    setState(() {
      _selectedTab = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          primary: true,
          child: Padding(
            padding: EdgeInsets.all(6.w),
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
                  onChanged: (String text) => setState(() => query = text),
                  padding: EdgeInsets.symmetric(horizontal: 4.w, vertical: 3.w),
                  autofocus: true,
                  placeholder: "What are you looking for?",
                ),
                SizedBox(height: 2.w),
                Row(
                  children: [
                    _Tab(
                      text: "words",
                      onSelect: () => selectTab(0),
                      active: _selectedTab != 0,
                    ),
                    SizedBox(width: 2.w),
                    _Tab(
                      text: "movies",
                      onSelect: () => selectTab(1),
                      active: _selectedTab != 1,
                    ),
                    // SizedBox(width: 2.w),
                    // _Tab(
                    //   text: "books",
                    //   onSelect: () => selectTab(2),
                    //   active: _selectedTab != 2,
                    // ),
                  ],
                ),
                SizedBox(height: 2.w),
                Stack(
                  children: [
                    Visibility(
                      visible: _selectedTab == 0,
                      maintainState: true,
                      child: SearchList<WordEntry>(
                        query: query,
                        future: dictionary_api.search(query),
                        itemBuilder: (entry) => WordSearchItem(entry: entry),
                      ),
                    ),
                    Visibility(
                      visible: _selectedTab == 1,
                      maintainState: true,
                      child: SearchList<Movie>(
                        query: query,
                        future: movies_api.search(query),
                        itemBuilder: (movie) => MovieSearchItem(movie: movie),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

class SearchList<T> extends StatelessWidget {
  final String query;
  final Future<List<T>> future;
  final Widget Function(T) itemBuilder;

  const SearchList({
    Key? key,
    required this.query,
    required this.future,
    required this.itemBuilder,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<T>>(
      future: future,
      builder: (context, snapshot) {
        if (!snapshot.hasData) return const SizedBox.shrink();

        final results = snapshot.data!;
        return ListView.separated(
          physics: const NeverScrollableScrollPhysics(),
          itemBuilder: (context, index) => itemBuilder(results[index]),
          separatorBuilder: (context, index) => SizedBox(height: 2.w),
          itemCount: results.length,
          shrinkWrap: true,
        );
      },
    );
  }
}

class _Tab extends StatelessWidget {
  final String text;
  final VoidCallback onSelect;
  final bool active;

  const _Tab({
    Key? key,
    required this.text,
    required this.onSelect,
    required this.active,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 25.w,
      height: 9.w,
      child: ElevatedButton(
        onPressed: onSelect,
        style: ElevatedButton.styleFrom(
          minimumSize: Size.zero,
          elevation: 0,
          primary: active
              ? const Color.fromRGBO(23, 23, 23, 1)
              : const Color.fromRGBO(54, 54, 54, 1),
          onPrimary: const Color.fromRGBO(54, 54, 54, 1),
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
