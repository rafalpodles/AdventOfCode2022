#!/usr/bin/env python
'''
This script will grab the leaderboard from Advent of Code and post it to Slack
'''
# pylint: disable=wrong-import-order
# pylint: disable=C0301,C0103,C0209

import datetime
import sys
import json
import os
import secrets
from collections import defaultdict

import requests
import schedule

LEADERBOARD_ID = secrets.LEADERBOARD_ID
SESSION_ID = secrets.SESSION_ID
SLACK_WEBHOOK = secrets.SLACK_WEBHOOK

# You should not need to change this URL
LEADERBOARD_URL = "https://adventofcode.com/{}/leaderboard/private/view/{}".format(
        datetime.datetime.today().year,
        LEADERBOARD_ID)


def format_leader_message(members):
    """
    Format the message to conform to Slack's API
    """
    message = ""

    # add each member to message
    medals = [':third_place_medal:', ':second_place_medal:', ':trophy:']
    for username, score, stars in members:
        if medals:
            medal = ' ' + medals.pop()
        else:
            medal = ''
        message += f"{medal}*{username}* {score} Points, {stars} Stars\n"

    message += f"\n<{LEADERBOARD_URL}|View Leaderboard Online>"

    return message


def parse_members(members_json):
    """
    Handle member lists from AoC leaderboard
    """
    members = [(m["name"],
                m["local_score"],
                m["stars"]
                ) for m in members_json.values()]

    # sort members by score, descending
    members.sort(key=lambda s: (-s[1], -s[2]))

    return members


def post_message(message):
    """
    Post the message to Slack's API in the proper channel
    """
    payload = json.dumps({
        "icon_emoji": ":christmas_tree:",
        "username": "Advent Of Code Leaderboard",
        "text": message
    })

    requests.post(
        SLACK_WEBHOOK,
        data=payload,
        timeout=60,
        headers={"Content-Type": "application/json"}
    )


def set_persons(persons_in_json):  # Needed to modify global copy of globvar
    text_file = open("persons.json", "w+")
    text_file.write(persons_in_json)
    text_file.close()


def get_persons():
    try:
        text_file = open("persons.json", "r")
        persons_in_json = text_file.read()
        text_file.close()
        return persons_in_json
    except FileNotFoundError:
        print("Not found persons file. Will return empty string")
        return ""


def job():
    # retrieve leaderboard
    print("Running cron")
    r = requests.get(
        "{}.json".format(LEADERBOARD_URL),
        timeout=60,
        cookies={"session": SESSION_ID}
    )

    if r.status_code != requests.codes.ok:  # pylint: disable=no-member
        print("Error retrieving leaderboard")
        sys.exit(1)

    members_json = r.json()["members"]
    update = False

    # person class
    class Person:
        def __init__(self, name, local_score, stars):
            self.name = name
            self.local_score = local_score
            self.stars = stars

    # extract persons from json
    last_persons = []
    try:
        j = json.loads(get_persons())
        for p in j:
            last_persons.append(Person(p["name"], p["local_score"], p["stars"]))
    except Exception:
        print("No persons in db")

    # get persons from response
    persons = [Person(m["name"],
                      m["local_score"],
                      m["stars"]
                      ) for m in members_json.values()]

    # check if there are changes in scores
    if len(persons) == len(last_persons):
        for lp in last_persons:
            for p in persons:
                if lp.name == p.name:
                    if lp.local_score != p.local_score or lp.stars != p.stars:
                        update = True
    else:
        update = True

    if update:
        print("Found changes. Updating...")
        members = parse_members(members_json)
        # generate message to send to slack
        message = format_leader_message(members)
        post_message(message)
    else:
        print("No changes.")

    # store persons from response to file
    set_persons(json.dumps([obj.__dict__ for obj in persons], ensure_ascii=False))

def main():
    # make sure all variables are filled
    if LEADERBOARD_ID == "" or SESSION_ID == "" or SLACK_WEBHOOK == "":
        print("Please update script variables before running script.\n\
                See README for details on how to do this.")
        sys.exit(1)
    # """
    # Main program loop
    # """
    print("Starting app")
    job()


if __name__ == "__main__":
    main()



