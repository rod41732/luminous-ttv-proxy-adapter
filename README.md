# Luminous TTV proxy adapter

This is simple service to proxy the incoming request to upstream Luminous TTV server
https://github.com/AlyoshaVasilieva/luminous-ttv

## Why this exist.

In iOS app, I'm using [Amazon IVS Player](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/player-ios.html) to play Twitch Streams.

The problem is when using URL like `https://uminous.alienpls.org/live/forsen?platform=web&allow_source=true&allow_audio_only=true&fast_bread=true&warp=true&supported_codecs=av1%2ch265%2ch264`, the player failed to play with error like this

```
[IVSPlayer] Amazon IVS Player SDK 1.35.0
[IVSPlayer] Player stopping playback - error Player:3 (ErrorNoSource code 0 - Source create failed)
```

Turns out, the player struggle to create source if there's no `.m3u8` extension in the path name

## Feature

- Convert url `/proxy/{channel}` or `/proxy/{channel}.m3u8` to `{host}/live/{channel}{queryParams}
- sensible default, if not specified host/params default to "a working host" and "query params that allow low latency stream" which is most likely the use case

## Usage

- Build and run this app
- Point the video player to `localhost:8080/proxy/{channel}.m3e8`

  - or try `curl` it, it should result in m3u8 response

  ```
  #EXTM3U
  #EXT-X-TWITCH-INFO:NODE="video-edge-dc44ae.jfk50",MANIFEST-NODE-TYPE="weaver_cluster",MANIFEST-NODE="video-weaver.jfk04",SUPPRESS="true",SERVER-TIME="1739692448.33",TRANSCODESTACK="2017TranscodeX264_V2",TRANSCODEMODE="cbr_v1",USER-IP="1.1.1.1",SERVING-ID="2bd1a4bb8839467ab33189a22e2b143d",CLUSTER="jfk50",ABS="true",VIDEO-SESSION-ID="8981770687066833309",BROADCAST-ID="316207780604",STREAM-TIME="44824.326398",FUTURE="true",B="false",USER-COUNTRY="US",MANIFEST-CLUSTER="jfk04",ORIGIN="pdx05",C="xxxxxxxx",D="false",E="xxxxxxxx"
  #EXT-X-MEDIA:TYPE=VIDEO,GROUP-ID="chunked",NAME="936p60 (source)",AUTOSELECT=YES,DEFAULT=YES

  #EXT-X-STREAM-INF:BANDWIDTH=9164743,RESOLUTION=1664x936,CODECS="avc1.64002A,mp4a.40.2",VIDEO="chunked",FRAME-RATE=60.000
  https://video-weaver.jfk04.hls.ttvnw.net/v1/playlist/xxxxx

  #EXT-X-MEDIA:TYPE=VIDEO,GROUP-ID="720p60",NAME="720p60",AUTOSELECT=YES,DEFAULT=YES
  #EXT-X-STREAM-INF:BANDWIDTH=3422999,RESOLUTION=1280x720,CODECS="avc1.4D401F,mp4a.40.2",VIDEO="720p60",FRAME-RATE=60.000
  https://video-weaver.jfk04.hls.ttvnw.net/v1/playlist/xxxxx

  #EXT-X-MEDIA:TYPE=VIDEO,GROUP-ID="720p30",NAME="720p",AUTOSELECT=YES,DEFAULT=YES
  #EXT-X-STREAM-INF:BANDWIDTH=2373000,RESOLUTION=1280x720,CODECS="avc1.4D401F,mp4a.40.2",VIDEO="720p30",FRAME-RATE=30.000
  https://video-weaver.jfk04.hls.ttvnw.net/v1/playlist/xxxx

  ```

- The video should start playing properly

## Notes

- If there's error like 404 not found, it's likely because the stream is not live, try with live channel

```

```
