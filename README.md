# 🎨 OsuStatsApi 🎶

Osu Stats Api provides fast and powerful
access to player statistics,
including real PP count and real Accuracy percentage

<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osu-stats-api.herokuapp.com/user/29829158
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
   "username":"hud0shnik",
   "previous_usernames":"",
   "avatar_url":"https://a.ppy.sh/29829158?1651433350.jpeg",
   "id":"29829158",
   "country_code":"RU",
   "global_rank":"875126",
   "country_rank":"77743",
   "pp":"616.145",
   "play_time":"21h18m38s",
   "play_time_seconds":"76718",
   "ssh":"101",
   "ss":"76",
   "sh":"0",
   "s":"3",
   "a":"6",
   "ranked_score":"84787365",
   "accuracy":"99.9987",
   "play_count":"1836",
   "total_score":"305266199",
   "total_hits":"89998",
   "maximum_combo":"724",
   "replays":"2",
   "level":"36",
   "support_level":"0",
   "default_group":"default",
   "is_online":"false",
   "is_active":"true",
   "is_deleted":"false",
   "is_bot":"false",
   "is_supporter":"false",
   "last_visit":"null",
   "profile_color":"null",
   "ranked_beatmapset_count":"0",
   "pending_beatmapset_count":"0",
   "pm_friends_only":"false",
   "graveyard_beatmapset_count":"0",
   "beatmap_playcounts_count":"287",
   "comments_count":"0",
   "favorite_beatmapset_count":"6",
   "guest_beatmapset_count":"0",
   "follower_count":"2",
   "best_beat_map":{
      "difficulty_rating":"1.59",
      "id":"1445261",
      "rank":"XH",
      "ended_at":"2022-05-14T18:33:06+00:00",
      "accuracy":"1",
      "maximum_combo":"65",
      "pp":"44.3059",
      "passed":"true",
      "total_score":"78652",
      "legacy_perfect":"true",
      "replay":"false",
      "mode":"osu",
      "status":"ranked",
      "total_length":"27",
      "ar":"5.2",
      "bpm":"76",
      "convert":"false",
      "count_circles":"8",
      "count_sliders":"27",
      "count_spinners":"1",
      "cs":"2.8",
      "deleted_at":"null",
      "drain":"3.6",
      "hit_length":"26",
      "is_scoreable":"true",
      "last_updated":"2017-11-07T03:57:30+00:00",
      "mode_int":"0",
      "pass_count":"40815",
      "play_count":"78318",
      "ranked":"1",
      "url":"https://osu.ppy.sh/beatmaps/1445261",
      "checksum":"3c8dee086cfd75224875253ed1f227b6",
      "creator":"fieryrage",
      "favorite_count":"99",
      "hype":"null",
      "nsfw":"false",
      "offset":"0",
      "spotlight":"false"
   }
   ```
