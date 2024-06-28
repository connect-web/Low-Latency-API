import os

bosses = {
    "Bounty Hunter - Hunter": {},
    "Bounty Hunter - Rogue": {},
    "Bounty Hunter (Legacy) - Hunter": {},
    "Bounty Hunter (Legacy) - Rogue": {},
    "Clue Scrolls (all)": {},
    "Clue Scrolls (beginner)": {},
    "Clue Scrolls (easy)": {},
    "Clue Scrolls (medium)": {},
    "Clue Scrolls (hard)": {},
    "Clue Scrolls (elite)": {},
    "Clue Scrolls (master)": {},
    "LMS - Rank": {},
    "PvP Arena - Rank": {},
    "Soul Wars Zeal": {},
    "Rifts closed": {},
    "Abyssal Sire": {},
    "Alchemical Hydra": {},
    "Artio": {},
    "Barrows Chests": {},
    "Bryophyta": {},
    "Callisto": {},
    "Calvar'ion": {},
    "Cerberus": {},
    "Chambers of Xeric": {},
    "Chambers of Xeric: Challenge Mode": {},
    "Chaos Elemental": {},
    "Chaos Fanatic": {},
    "Commander Zilyana": {},
    "Corporeal Beast": {},
    "Crazy Archaeologist": {},
    "Dagannoth Prime": {},
    "Dagannoth Rex": {},
    "Dagannoth Supreme": {},
    "Deranged Archaeologist": {},
    "Duke Sucellus": {},
    "General Graardor": {},
    "Giant Mole": {},
    "Grotesque Guardians": {},
    "Hespori": {},
    "Kalphite Queen": {},
    "King Black Dragon": {},
    "Kraken": {},
    "Kree'Arra": {},
    "K'ril Tsutsaroth": {},
    "Mimic": {},
    "Nex": {},
    "Nightmare": {},
    "Phosani's Nightmare": {},
    "Obor": {},
    "Phantom Muspah": {},
    "Sarachnis": {},
    "Scorpia": {},
    "Skotizo": {},
    "Spindel": {},
    "Tempoross": {},
    "The Gauntlet": {},
    "The Corrupted Gauntlet": {},
    "The Leviathan": {},
    "The Whisperer": {},
    "Theatre of Blood": {},
    "Theatre of Blood: Hard Mode": {},
    "Thermonuclear Smoke Devil": {},
    "Tombs of Amascut": {},
    "Tombs of Amascut: Expert Mode": {},
    "TzKal-Zuk": {},
    "TzTok-Jad": {},
    "Vardorvis": {},
    "Venenatis": {},
    "Vet'ion": {},
    "Vorkath": {},
    "Wintertodt": {},
    "Zalcano": {},
    "Zulrah": {},
    "Deadman Points": {},
    "League Points": {},
    "Colosseum Glory": {},
}

activities = {
    "Soul Wars Zeal" : "/activities/soul_wars_zeal",
    "Rifts closed" : "/activities/rifts_closed",
    "LMS - Rank" : "/activities/last_man_standing",
    "Colosseum Glory" : "/activities/colosseum_glory",
    "Clue Scrolls (all)" : "/activities/clue_scroll_all",
    "Clue Scrolls (beginner)" : "/activities/external/clue_scroll_beginner",
    "Clue Scrolls (easy)" : "/activities/external/clue_scroll_easy",
    "Clue Scrolls (medium)" : "/activities/external/clue_scroll_medium",
    "Clue Scrolls (hard)" : "/activities/external/clue_scroll_hard",
    "Clue Scrolls (elite)" : "/activities/external/clue_scroll_elite",
    "Clue Scrolls (master)" : "/activities/external/clue_scroll_master",

    "Bounty Hunter - Hunter" : "/activities/bounty_hunter_hunter",
    "Bounty Hunter - Rogue" : "/activities/bounty_hunter_rogue",
    "Bounty Hunter (Legacy) - Hunter" : "/activities/bounty_hunter_hunter",
    "Bounty Hunter (Legacy) - Rogue" : "/activities/bounty_hunter_rogue",
    "League Points" : "/activities/league_points",
    "PvP Arena - Rank" : "/activities/pvp_arena_rank",


"Tombs of Amascut: Expert Mode" : "/bosses/tombs_of_amascut_expert",
    "TzKal-Zuk": "/bosses/tzkal_zuk",
    "TzTok-Jad": "/bosses/tztok_jad",
    "Deadman Points": "/deadman",
}

import re


def filterFileName(input):
    # Use regex to replace all non a-zA-Z characters, but keep spaces
    output = re.sub(r'[^a-zA-Z\s]', '', input)
    output = output.lower()
    output = output.replace(' ', '_')
    return output

file_names = {}

for file in os.listdir('bosses'):
    file = file.split('.')[0]
    file_names[file] = True



for boss in bosses:
    filtered_name = filterFileName(boss)

    if filtered_name not in file_names:
        if boss in activities:
            path = '.'+activities[boss]+'.png'
            if not os.path.exists(path):
                print(path) # print path if it does not exist
        else:
            print(boss) # print if boss not in activities map

    else:
        # if the boss does not exist then print the path to fix.
        path = f'./bosses/{filtered_name}.png'
        if not os.path.exists(path):
            print(path)

