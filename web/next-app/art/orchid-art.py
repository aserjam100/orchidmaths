"""Draws Orchid's footer and login artwork as SVG. Original work, MIT like the rest of the repo.

Run `make art` from the repo root: this writes art/out/*.svg, and rsvg-convert renders each
to the PNG the app ships. Deterministic: a fixed seed per picture, so a re-run reproduces
the same files.
"""
import math
import os
import random

OUT = os.path.join(os.path.dirname(__file__), "out")
os.makedirs(OUT, exist_ok=True)

PETAL = ["#c86fa8", "#d98cbc", "#b45a96", "#e4a9cf", "#9d4c88"]
LEAF = "#5f8a6a"


def flower(x, y, s, rot, rng):
    c = rng.choice(PETAL)
    lip = rng.choice(PETAL)
    g = [f'<g transform="translate({x:.1f},{y:.1f}) rotate({rot:.1f}) scale({s:.2f})" opacity="0.92">']
    for a in (0, 72, 144, 216, 288):  # three sepals and two petals
        g.append(f'<ellipse cx="0" cy="-22" rx="9" ry="22" fill="{c}" transform="rotate({a})"/>')
    g.append(f'<path d="M-10,4 C-14,22 14,22 10,4 Z" fill="{lip}" stroke="#7a3568" stroke-width="1"/>')
    g.append('<circle cx="0" cy="0" r="4.5" fill="#f6e27a"/>')
    g.append("</g>")
    return "".join(g)


def spray(x0, y0, length, ang, n, scale, rng):
    """An arching stem with flowers along its outer two thirds."""
    steps = n * 3
    pts = []
    for i in range(steps + 1):
        t = i / steps
        a = math.radians(ang - 70 * t * (1 if ang < 90 else -1))
        pts.append((x0 + length * t * math.cos(a), y0 - length * t * math.sin(a)))
    d = "M" + " L".join(f"{px:.1f},{py:.1f}" for px, py in pts)
    out = [f'<path d="{d}" fill="none" stroke="{LEAF}" stroke-width="{3 * scale:.1f}" stroke-linecap="round"/>']
    for i in range(n):
        px, py = pts[steps // 3 + i * (2 * steps // 3) // n]
        out.append(flower(px, py, scale * rng.uniform(0.8, 1.15), rng.uniform(-30, 30), rng))
    return "".join(out)


def leaves(x, y, s):
    return (
        f'<path d="M{x},{y} C{x - 140 * s},{y - 40 * s} {x - 200 * s},{y - 10 * s} {x - 230 * s},{y + 5 * s} '
        f'C{x - 150 * s},{y + 15 * s} {x - 60 * s},{y + 10 * s} {x},{y} Z" fill="{LEAF}" opacity="0.8"/>'
        f'<path d="M{x},{y} C{x + 120 * s},{y - 50 * s} {x + 190 * s},{y - 30 * s} {x + 220 * s},{y - 15 * s} '
        f'C{x + 140 * s},{y + 5 * s} {x + 60 * s},{y + 12 * s} {x},{y} Z" fill="#6f9c78" opacity="0.8"/>'
    )


def paper(w, h, step, rng):
    """Faint graph paper and a scatter of symbols: the maths under the flowers."""
    g = ['<g stroke="#b9a6c4" stroke-width="1" opacity="0.35">']
    g += [f'<line x1="{x}" y1="0" x2="{x}" y2="{h}"/>' for x in range(0, w, step)]
    g += [f'<line x1="0" y1="{y}" x2="{w}" y2="{y}"/>' for y in range(0, h, step)]
    g.append("</g>")
    syms = ["½", "×", "÷", "π", "+", "=", "%", "¾", "√", "△"]
    for _ in range(int(w * h / 60000)):
        g.append(
            f'<text x="{rng.uniform(0, w):.0f}" y="{rng.uniform(0, h):.0f}" font-family="Georgia,serif" '
            f'font-size="{rng.uniform(18, 34):.0f}" fill="#8e6fa3" opacity="0.28">{rng.choice(syms)}</text>'
        )
    return "".join(g)


def picture(name, w, h, seed, sprays):
    rng = random.Random(seed)
    body = [
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}">',
        '<defs><linearGradient id="bg" x1="0" y1="0" x2="0" y2="1">'
        '<stop offset="0" stop-color="#fbf7fc"/><stop offset="1" stop-color="#efe3f3"/></linearGradient></defs>',
        f'<rect width="{w}" height="{h}" fill="url(#bg)"/>',
        paper(w, h, 48, rng),
    ]
    for fx, fy, length, ang, n, sc in sprays:
        x, y = fx * w, fy * h
        body.append(leaves(x, y, sc))
        body.append(spray(x, y, length * min(w, h), ang, n, sc, rng))
    body.append("</svg>")
    with open(os.path.join(OUT, name + ".svg"), "w") as f:
        f.write("".join(body))


# Footer, desktop: a wide band kept pale so the footer's dark ink reads on it (tone "light").
picture("footer-orchid-wide", 1672, 643, 1, [(0.06, 1.0, 0.9, 80, 5, 1.1), (0.94, 1.0, 0.85, 100, 4, 1.0)])
picture("footer-orchid-wide-2", 1672, 643, 2, [(0.1, 1.0, 0.95, 70, 5, 1.0), (0.88, 1.02, 0.8, 110, 4, 1.15)])
# Footer, phone portraits.
picture("footer-orchid-tall", 1080, 1920, 3, [(0.12, 1.0, 0.9, 80, 6, 1.6)])
picture("footer-orchid-tall-2", 1080, 1920, 4, [(0.88, 1.0, 0.9, 100, 6, 1.6)])
# Login hero, portrait.
picture("login-orchid", 736, 1051, 5, [(0.2, 1.0, 0.95, 78, 6, 1.3), (0.85, 1.02, 0.7, 102, 3, 1.0)])
