/*!
 *
 * js-audio-recorder - js audio recorder plugin
 *
 * @version v0.5.7
 * @homepage https://github.com/2fps/recorder
 * @author 2fps <echoweb@126.com> (https://www.zhuyuntao.cn)
 * @license MIT
 *
 */
!(function (t, e) {
  "object" == typeof exports && "object" == typeof module
    ? (module.exports = e())
    : "function" == typeof define && define.amd
    ? define([], e)
    : "object" == typeof exports
    ? (exports.Recorder = e())
    : (t.Recorder = e());
})(this, function () {
  return (function (t) {
    var e = {};

    function i(n) {
      if (e[n]) return e[n].exports;
      var r = (e[n] = {
        i: n,
        l: !1,
        exports: {},
      });
      return t[n].call(r.exports, r, r.exports, i), (r.l = !0), r.exports;
    }

    return (
      (i.m = t),
      (i.c = e),
      (i.d = function (t, e, n) {
        i.o(t, e) ||
          Object.defineProperty(t, e, {
            enumerable: !0,
            get: n,
          });
      }),
      (i.r = function (t) {
        "undefined" != typeof Symbol &&
          Symbol.toStringTag &&
          Object.defineProperty(t, Symbol.toStringTag, {
            value: "Module",
          }),
          Object.defineProperty(t, "__esModule", {
            value: !0,
          });
      }),
      (i.t = function (t, e) {
        if ((1 & e && (t = i(t)), 8 & e)) return t;
        if (4 & e && "object" == typeof t && t && t.__esModule) return t;
        var n = Object.create(null);
        if (
          (i.r(n),
          Object.defineProperty(n, "default", {
            enumerable: !0,
            value: t,
          }),
          2 & e && "string" != typeof t)
        )
          for (var r in t)
            i.d(
              n,
              r,
              function (e) {
                return t[e];
              }.bind(null, r)
            );
        return n;
      }),
      (i.n = function (t) {
        var e =
          t && t.__esModule
            ? function () {
                return t.default;
              }
            : function () {
                return t;
              };
        return i.d(e, "a", e), e;
      }),
      (i.o = function (t, e) {
        return Object.prototype.hasOwnProperty.call(t, e);
      }),
      (i.p = ""),
      i((i.s = 0))
    );
  })([
    function (t, e, i) {
      "use strict";
      Object.defineProperty(e, "__esModule", {
        value: !0,
      });
      var n = (function () {
        function t(e) {
          void 0 === e && (e = {}),
            (this.isplaying = !1),
            (this.lBuffer = []),
            (this.rBuffer = []),
            (this.tempPCM = []),
            (this.inputSampleBits = 16),
            (this.playStamp = 0),
            (this.playTime = 0),
            (this.totalPlayTime = 0),
            (this.offset = 0),
            (this.fileSize = 0);
          var i,
            n = new (window.AudioContext || window.webkitAudioContext)();
          (this.inputSampleRate = n.sampleRate),
            (this.config = {
              sampleBits: ~[8, 16].indexOf(e.sampleBits) ? e.sampleBits : 16,
              sampleRate: ~[8e3, 11025, 16e3, 22050, 24e3, 44100, 48e3].indexOf(
                e.sampleRate
              )
                ? e.sampleRate
                : this.inputSampleRate,
              numChannels: ~[1, 2].indexOf(e.numChannels) ? e.numChannels : 1,
              compiling: !!e.compiling || !1,
            }),
            (this.outputSampleRate = this.config.sampleRate),
            (this.oututSampleBits = this.config.sampleBits),
            (this.littleEdian =
              ((i = new ArrayBuffer(2)),
              new DataView(i).setInt16(0, 256, !0),
              256 === new Int16Array(i)[0])),
            t.initUserMedia();
        }

        return (
          (t.prototype.initRecorder = function () {
            var t = this;
            this.context && this.destroy(),
              (this.context = new (window.AudioContext ||
                window.webkitAudioContext)()),
              (this.analyser = this.context.createAnalyser()),
              (this.analyser.fftSize = 2048);
            var e =
              this.context.createScriptProcessor ||
              this.context.createJavaScriptNode;
            (this.recorder = e.apply(this.context, [
              4096,
              this.config.numChannels,
              this.config.numChannels,
            ])),
              (this.recorder.onaudioprocess = function (e) {
                if (t.isrecording && !t.ispause) {
                  var i,
                    n = e.inputBuffer.getChannelData(0),
                    r = null;
                  if (
                    (t.lBuffer.push(new Float32Array(n)),
                    (t.size += n.length),
                    2 === t.config.numChannels &&
                      ((r = e.inputBuffer.getChannelData(1)),
                      t.rBuffer.push(new Float32Array(r)),
                      (t.size += r.length)),
                    t.config.compiling)
                  ) {
                    var o = t.transformIntoPCM(n, r);
                    t.tempPCM.push(o),
                      (t.fileSize = o.byteLength * t.tempPCM.length);
                  } else
                    t.fileSize =
                      Math.floor(
                        t.size /
                          Math.max(t.inputSampleRate / t.outputSampleRate, 1)
                      ) *
                      (t.oututSampleBits / 8);
                  (i = 100 * Math.max.apply(Math, n)),
                    (t.duration += 4096 / t.inputSampleRate),
                    t.onprocess && t.onprocess(t.duration),
                    t.onprogress &&
                      t.onprogress({
                        duration: t.duration,
                        fileSize: t.fileSize,
                        vol: i,
                        data: t.tempPCM,
                      });
                }
              });
          }),
          (t.prototype.start = function () {
            var t = this;
            if (!this.isrecording)
              return (
                this.clear(),
                this.initRecorder(),
                (this.isrecording = !0),
                navigator.mediaDevices
                  .getUserMedia({
                    audio: !0,
                  })
                  .then(function (e) {
                    (t.audioInput = t.context.createMediaStreamSource(e)),
                      (t.stream = e);
                  })
                  .then(function () {
                    t.audioInput.connect(t.analyser),
                      t.analyser.connect(t.recorder),
                      t.recorder.connect(t.context.destination);
                  })
              );
          }),
          (t.prototype.pause = function () {
            this.isrecording && !this.ispause && (this.ispause = !0);
          }),
          (t.prototype.resume = function () {
            this.isrecording && this.ispause && (this.ispause = !1);
          }),
          (t.prototype.stop = function () {
            (this.isrecording = !1),
              this.audioInput && this.audioInput.disconnect(),
              this.recorder.disconnect();
          }),
          (t.prototype.play = function () {
            this.stop(),
              this.source && this.source.stop(),
              (this.isplaying = !0),
              (this.playTime = 0),
              this.playAudioData();
          }),
          (t.prototype.getPlayTime = function () {
            var t = 0;
            return (
              (t = this.isplaying
                ? this.context.currentTime - this.playStamp + this.playTime
                : this.playTime) >= this.totalPlayTime &&
                (t = this.totalPlayTime),
              t
            );
          }),
          (t.prototype.pausePlay = function () {
            !this.isrecording &&
              this.isplaying &&
              (this.source && this.source.disconnect(),
              (this.playTime += this.context.currentTime - this.playStamp),
              (this.isplaying = !1));
          }),
          (t.prototype.resumePlay = function () {
            this.isrecording ||
              this.isplaying ||
              0 === this.playTime ||
              ((this.isplaying = !0), this.playAudioData());
          }),
          (t.prototype.stopPlay = function () {
            this.isrecording ||
              ((this.playTime = 0),
              (this.isplaying = !1),
              this.source && this.source.stop());
          }),
          (t.prototype.getWholeData = function () {
            return this.tempPCM;
          }),
          (t.prototype.getNextData = function () {
            var t1 = this.tempPCM.length,
              e1 = this.tempPCM.slice(this.offset);

            if (e1.length) {
              var e = new ArrayBuffer(e1.length * e1[0].byteLength),
                i = new DataView(e),
                n = 0;
              e1.forEach(function (t) {
                for (var e = 0, r = t.byteLength; e < r; ++e)
                  i.setInt8(n, t.getInt8(e)), n++;
              }),
                (this.PCM = i),
                (this.tempPCM = []);
            }
            if (this.PCM) return this.PCM;
            var r = this.flat();
            (r = t.compress(r, this.inputSampleRate, this.outputSampleRate)),
              (this.PCM = t.encodePCM(
                r,
                this.oututSampleBits,
                this.littleEdian
              ));
            this.offset = t1;
            return new Blob([r]);
            // return  e1
          }),
          (t.prototype.playAudioData = function () {
            var e = this;
            this.context.decodeAudioData(
              this.getWAV().buffer,
              function (t) {
                (e.source = e.context.createBufferSource()),
                  (e.source.buffer = t),
                  (e.totalPlayTime = e.source.buffer.duration),
                  e.source.connect(e.analyser),
                  e.analyser.connect(e.context.destination),
                  e.source.start(0, e.playTime),
                  (e.playStamp = e.context.currentTime);
              },
              function (e) {
                t.throwError(e);
              }
            );
          }),
          (t.prototype.getRecordAnalyseData = function () {
            if (this.ispause) return this.prevDomainData;
            var t = new Uint8Array(this.analyser.frequencyBinCount);
            return (
              this.analyser.getByteTimeDomainData(t), (this.prevDomainData = t)
            );
          }),
          (t.prototype.getPlayAnalyseData = function () {
            return this.getRecordAnalyseData();
          }),
          (t.prototype.getPCM = function () {
            if (this.tempPCM.length) {
              var e = new ArrayBuffer(
                  this.tempPCM.length * this.tempPCM[0].byteLength
                ),
                i = new DataView(e),
                n = 0;
              this.tempPCM.forEach(function (t) {
                for (var e = 0, r = t.byteLength; e < r; ++e)
                  i.setInt8(n, t.getInt8(e)), n++;
              }),
                (this.PCM = i),
                (this.tempPCM = []);
            }
            if (this.PCM) return this.PCM;
            var r = this.flat();
            return (
              (r = t.compress(r, this.inputSampleRate, this.outputSampleRate)),
              (this.PCM = t.encodePCM(
                r,
                this.oututSampleBits,
                this.littleEdian
              ))
            );
          }),
          (t.prototype.getPCMBlob = function () {
            return this.stop(), new Blob([this.getPCM()]);
          }),
          (t.prototype.downloadPCM = function (t) {
            void 0 === t && (t = "recorder");
            var e = this.getPCMBlob();
            this.download(e, t, "pcm");
          }),
          (t.prototype.getWAV = function () {
            var e = this.getPCM();
            return t.encodeWAV(
              e,
              this.inputSampleRate,
              this.outputSampleRate,
              this.config.numChannels,
              this.oututSampleBits,
              this.littleEdian
            );
          }),
          (t.prototype.getWAVBlob = function () {
            return (
              this.stop(),
              new Blob([this.getWAV()], {
                type: "audio/wav",
              })
            );
          }),
          (t.prototype.downloadWAV = function (t) {
            void 0 === t && (t = "recorder");
            var e = this.getWAVBlob();
            this.download(e, t, "wav");
          }),
          (t.prototype.transformIntoPCM = function (e, i) {
            var n = new Float32Array(e),
              r = new Float32Array(i),
              o = t.compress(
                {
                  left: n,
                  right: r,
                },
                this.inputSampleRate,
                this.outputSampleRate
              );
            return t.encodePCM(o, this.oututSampleBits, this.littleEdian);
          }),
          (t.prototype.destroy = function () {
            return this.stopStream(), this.closeAudioContext();
          }),
          (t.prototype.stopStream = function () {
            this.stream &&
              this.stream.getTracks &&
              (this.stream.getTracks().forEach(function (t) {
                return t.stop();
              }),
              (this.stream = null));
          }),
          (t.prototype.closeAudioContext = function () {
            return this.context &&
              this.context.close &&
              "closed" !== this.context.state
              ? this.context.close()
              : new Promise(function (t) {
                  t();
                });
          }),
          (t.prototype.download = function (e, i, n) {
            try {
              var r = document.createElement("a");
              (r.href = window.URL.createObjectURL(e)),
                (r.download = i + "." + n),
                r.click();
            } catch (e) {
              t.throwError(e);
            }
          }),
          (t.prototype.clear = function () {
            (this.lBuffer.length = 0),
              (this.rBuffer.length = 0),
              (this.size = 0),
              (this.fileSize = 0),
              (this.PCM = null),
              (this.audioInput = null),
              (this.duration = 0),
              (this.ispause = !1),
              (this.isplaying = !1),
              (this.playTime = 0),
              (this.totalPlayTime = 0),
              this.source && (this.source.stop(), (this.source = null));
          }),
          (t.prototype.flat = function () {
            var t = null,
              e = new Float32Array(0);
            1 === this.config.numChannels
              ? (t = new Float32Array(this.size))
              : ((t = new Float32Array(this.size / 2)),
                (e = new Float32Array(this.size / 2)));
            for (var i = 0, n = 0; n < this.lBuffer.length; n++)
              t.set(this.lBuffer[n], i), (i += this.lBuffer[n].length);
            i = 0;
            for (n = 0; n < this.rBuffer.length; n++)
              e.set(this.rBuffer[n], i), (i += this.rBuffer[n].length);
            return {
              left: t,
              right: e,
            };
          }),
          (t.playAudio = function (t) {
            var e = document.createElement("audio");
            (e.src = window.URL.createObjectURL(t)), e.play();
          }),
          (t.compress = function (t, e, i) {
            for (
              var n = e / i,
                r = Math.max(n, 1),
                o = t.left,
                s = t.right,
                a = Math.floor((o.length + s.length) / n),
                u = new Float32Array(a),
                h = 0,
                c = 0;
              h < a;

            ) {
              var l = Math.floor(c);
              (u[h] = o[l]), h++, s.length && ((u[h] = s[l]), h++), (c += r);
            }
            return u;
          }),
          (t.encodePCM = function (t, e, i) {
            void 0 === i && (i = !0);
            var n = 0,
              r = t.length * (e / 8),
              o = new ArrayBuffer(r),
              s = new DataView(o);
            if (8 === e)
              for (var a = 0; a < t.length; a++, n++) {
                var u =
                  (h = Math.max(-1, Math.min(1, t[a]))) < 0 ? 128 * h : 127 * h;
                (u = +u + 128), s.setInt8(n, u);
              }
            else
              for (a = 0; a < t.length; a++, n += 2) {
                var h = Math.max(-1, Math.min(1, t[a]));
                s.setInt16(n, h < 0 ? 32768 * h : 32767 * h, i);
              }
            return s;
          }),
          (t.encodeWAV = function (t, e, i, n, o, s) {
            void 0 === s && (s = !0);
            var a = i > e ? e : i,
              u = o,
              h = new ArrayBuffer(44 + t.byteLength),
              c = new DataView(h),
              l = n,
              p = 0;
            r(c, p, "RIFF"),
              (p += 4),
              c.setUint32(p, 36 + t.byteLength, s),
              r(c, (p += 4), "WAVE"),
              r(c, (p += 4), "fmt "),
              (p += 4),
              c.setUint32(p, 16, s),
              (p += 4),
              c.setUint16(p, 1, s),
              (p += 2),
              c.setUint16(p, l, s),
              (p += 2),
              c.setUint32(p, a, s),
              (p += 4),
              c.setUint32(p, l * a * (u / 8), s),
              (p += 4),
              c.setUint16(p, l * (u / 8), s),
              (p += 2),
              c.setUint16(p, u, s),
              r(c, (p += 2), "data"),
              (p += 4),
              c.setUint32(p, t.byteLength, s),
              (p += 4);
            for (var f = 0; f < t.byteLength; )
              c.setUint8(p, t.getUint8(f)), p++, f++;
            return c;
          }),
          (t.throwError = function (t) {
            throw new Error(t);
          }),
          (t.initUserMedia = function () {
            void 0 === navigator.mediaDevices && (navigator.mediaDevices = {}),
              void 0 === navigator.mediaDevices.getUserMedia &&
                (navigator.mediaDevices.getUserMedia = function (t) {
                  var e =
                    navigator.getUserMedia ||
                    navigator.webkitGetUserMedia ||
                    navigator.mozGetUserMedia;
                  return e
                    ? new Promise(function (i, n) {
                        e.call(navigator, t, i, n);
                      })
                    : Promise.reject(new Error("浏览器不支持 getUserMedia !"));
                });
          }),
          (t.getPermission = function () {
            return (
              this.initUserMedia(),
              navigator.mediaDevices
                .getUserMedia({
                  audio: !0,
                })
                .then(function (t) {
                  t.getTracks().forEach(function (t) {
                    return t.stop();
                  });
                })
            );
          }),
          t
        );
      })();

      function r(t, e, i) {
        for (var n = 0; n < i.length; n++) t.setUint8(e + n, i.charCodeAt(n));
      }

      e.default = n;
    },
  ]).default;
});
//# sourceMappingURL=recorder.js.map
