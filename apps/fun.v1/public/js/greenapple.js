var svgPathApple = "M376.349,171.58 c0,0-37.276,22.484-36.094,60.946s31.438,61.577,42.012,63.313c4.127,0.678-24.314,57.988-42.039,72.189 s-36.067,17.159-64.47,5.917c-28.403-11.242-48.521,0.724-65.089,6.871s-36.687-0.361-63.905-39.415 s-57.396-129.585-15.976-173.964s87.574-26.627,100-20.71s34.32,5.325,59.172-5.917S363.922,153.237,376.349,171.58z";
var svgPathLeaf = "M311.852,68.621c0,0,2.367,14.793-3.55,27.219 s-28.189,55.061-60.473,47.337c-0.809-0.193-5.529-14.482,1.398-29.002C259.004,93.682,284.49,70.699,311.852,68.621z";

var c = document.getElementById("c");
var ctx = c.getContext("2d");
var cw = c.width = 500,
  cx = cw / 2;
var ch = c.height = 500,
  cy = ch / 2;
var frames = 0;
var D = 25;
var colors = [70, 95, 120, 145, 170];
ctx.fillStyle = "rgba(0,0,0,.05)";

var Ry = [{
  p: new Path2D(svgPathApple),
  ry: new Array
}, {
  p: new Path2D(svgPathLeaf),
  ry: new Array
}, ]

function Particle(x, y) {
  this.x = x //cw/2 + r * Math.cos(a);
  this.y = y //ch/2 + r * Math.sin(a);
  this.ix = (Math.random()) * (Math.random() < 0.5 ? -1 : 1); //positive or negative
  this.iy = (Math.random()) * (Math.random() < 0.5 ? -1 : 1); //positive or negative
  this.hue = colors[Math.round(Math.random() * colors.length) + 1]
}

function createParticle(x, y, p) {
  var particle = new Particle(x, y);
  p.push(particle);
}

for (var k = 0; k < Ry.length; k++) {
  for (var j = 0; j < ch; j += 14) {
    for (var i = 0; i < cw; i += 14) {
      var x = i % cw;
      var y = j;
      var path = Ry[k].p;
      var p = Ry[k].ry;
      if (ctx.isPointInPath(path, x, y)) {
        createParticle(x, y, p);
      }
    }
  }
}

function Draw() {
  ctx.fillRect(0, 0, cw, ch);
  for (var i = 0; i < Ry.length; i++) {
    ctx.strokeStyle = "hsla(120,50%,20%,.5)";
    ctx.stroke(Ry[i].p);
    updateRy(Ry[i].ry, Ry[i].p);
    compare(Ry[i].ry);
  }
  window.requestAnimationFrame(Draw);
}

window.requestAnimationFrame(Draw);

function compare(p) {
  for (var i = 0; i < p.length; i++) {
    var a = p[i];
    for (var j = i + 1; j < p.length; j++) {

      var b = p[j];
      var dist = distance(a, b);
      if (dist < D) {
        var c = {};
        var alp = (D - dist) / D;
        var hue = a.hue;
        ctx.strokeStyle = "hsla(" + hue + ",87%, 44%," + alp + ")";
        ctx.beginPath();
        ctx.moveTo(a.x, a.y);
        ctx.lineTo(b.x, b.y);
        ctx.stroke();
      }
    }
  }
}

function updateRy(p, path) {
  //p : points array
  for (var i = 0; i < p.length; i++) {
    ctx.fillStyle = p[i].c;
    if (ctx.isPointInPath(path, p[i].x, p[i].y)) {
      p[i].x += p[i].ix;
      p[i].y += p[i].iy;

    } else {
      p[i].ix = -1 * p[i].ix;
      p[i].iy = -1 * p[i].iy;
      p[i].x += p[i].ix;
      p[i].y += p[i].iy;
    }
  }
  return p;
}

function distance(a, b) {
  var ac = b.y - a.y;
  var bc = b.x - a.x;
  var ab = Math.sqrt(ac * ac + bc * bc);
  return ab;
}