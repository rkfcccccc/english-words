import 'package:english_words/utils/api/api.dart';

import 'models.dart';

Future<Challenge> getChallenge() async {
  final data = await request(
    "get",
    "/vocabulary/challenge",
    authRequired: true,
  );

  return Challenge.fromMap(data);
}

Future<void> finishChallenge(String action, wordId) async {
  await request(
    "post",
    "/vocabulary/challenge",
    authRequired: true,
    body: {"action": action, "word_id": wordId},
  );
}

Future<void> setAlreadyLearned(String wordId, bool state) async {
  await request(
    "post",
    "/vocabulary/learned",
    authRequired: true,
    body: {"word_id": wordId, "state": state},
  );
}
