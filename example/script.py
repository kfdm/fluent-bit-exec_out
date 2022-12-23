import argparse
import json
import logging
import sys

parser = argparse.ArgumentParser()
parser.add_argument(
    "infile",
    type=argparse.FileType("r"),
    default=sys.stdin,
    nargs="?",
)
args = parser.parse_args()

logger = logging.getLevelName(__name__)

for line in args.infile:
    data = json.loads(line)
    logging.warning("from script %s", data)
