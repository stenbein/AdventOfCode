#!/usr/bin/python3
'''Day 11 of the 2017 advent of code'''

import pytest

from src.main import HexCounter


def helper(coords):
    '''helper function to run the test'''

    hexer = HexCounter()

    coords = coords.split(",")
    for coord in coords:
        hexer.move(coord)

    return hexer


def test_directions():
    '''test runner for some precalculated values'''

    #straight out
    hexer = helper("ne,ne,ne")
    assert hexer.max() == 3
    assert hexer.furthest == 3

    #there and back
    hexer = helper("ne,ne,sw,sw")
    assert hexer.max() == 0
    assert hexer.furthest == 2

    #angled back
    hexer = helper("ne,ne,s,s")
    assert hexer.max() == 2
    assert hexer.furthest == 2

    #zig zag
    hexer = helper("se,sw,se,sw,sw")
    assert hexer.max() == 3
    assert hexer.furthest == 3

    #loop
    hexer = helper("n,ne,ne,se,se,s,sw,sw,nw,nw")
    assert hexer.max() == 0
    assert hexer.furthest == 4
