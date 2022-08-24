import 'package:english_words/utils/api/api.dart';
import 'package:english_words/utils/api/dictionary/models.dart';

Future<List<WordEntry>> search(String query) async {
  final data = await request(
    "get",
    "/dictionary",
    params: {"query": query},
    authRequired: true,
  );

  return List<WordEntry>.from(data.map((e) => WordEntry.fromMap(e)));
}
