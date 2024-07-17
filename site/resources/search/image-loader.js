var bosses = {
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
    "TzKal-Zuk": {},
    "TzTok-Jad": {},
    "Vardorvis": {},
    "Venenatis": {},
    "Vet'ion": {},
    "Vorkath": {},
    "Wintertodt": {},
    "Zalcano": {},
    "Zulrah": {},
    "Scurrius": {},
}

var activities = {
    "Soul Wars Zeal": "/activities/soul_wars_zeal",
    "Rifts closed": "/activities/rifts_closed",
    "LMS - Rank": "/activities/last_man_standing",
    "Colosseum Glory": "/activities/colosseum_glory",
    "Clue Scrolls (all)": "/activities/clue_scroll_all",
    "Clue Scrolls (beginner)": "/activities/external/clue_scroll_beginner",
    "Clue Scrolls (easy)": "/activities/external/clue_scroll_easy",
    "Clue Scrolls (medium)": "/activities/external/clue_scroll_medium",
    "Clue Scrolls (hard)": "/activities/external/clue_scroll_hard",
    "Clue Scrolls (elite)": "/activities/external/clue_scroll_elite",
    "Clue Scrolls (master)": "/activities/external/clue_scroll_master",
    "Lunar Chests": "/activities/external/lunar_chest",


    "Bounty Hunter - Hunter": "/activities/bounty_hunter_hunter",
    "Bounty Hunter - Rogue": "/activities/bounty_hunter_rogue",
    "Bounty Hunter (Legacy) - Hunter": "/activities/bounty_hunter_hunter",
    "Bounty Hunter (Legacy) - Rogue": "/activities/bounty_hunter_rogue",
    "League Points": "/activities/league_points",
    "PvP Arena - Rank": "/activities/pvp_arena_rank",

    "Tombs of Amascut: Expert Mode": "/bosses/tombs_of_amascut_expert",
    "TzKal-Zuk": "/bosses/tzkal_zuk",
    "TzTok-Jad": "/bosses/tztok_jad",
    "Deadman Points": "/deadman",

};

function filterFileName(input) {
    // Use regex to replace all non a-zA-Z characters, but keep spaces
    let output = input.replace(/[^a-zA-Z\s]/g, '');
    // clean string to match filenames
    output = output.toLowerCase()
    output = output.replaceAll(' ', '_')
    return output
}


function getActivityImagePath(activity_name) {
    let filtered_activity_name = filterFileName(activity_name)
    if (activity_name in activities) {
        return `.${activities[activity_name]}.png`
    } else if (activity_name in bosses) {
        return `/bosses/${filtered_activity_name}.png`
    } else {
        console.log(`${activity_name} not in activities.`)
        return ""
    }

}


function titleCase(str) {
    return str.toLowerCase().replace(/(^|\s)\S/g, (letter) => letter.toUpperCase());
}

function viewLargeNumbers(value) {
    if (value > 10_000_000) {
        return `${(value / 1_000_000).toFixed(2)}M`;
    } else if (value > 100_000) {
        return `${(value / 1_000).toFixed(2)}K`;
    }
    return value.toString();
}

function getSkillImage(skill) {
    skill = skill.toLowerCase();
    return `<img src="/resources/imgs/hiscore/skill_icons/${skill}.png" alt="Icon ${skill}" class="img-fluid">`;
}

function getMinigameImage(minigame) {
    let minigame_path = getActivityImagePath(minigame)

    return `<img src="/resources/imgs/hiscore/${minigame_path}" alt="Icon ${minigame}" class="img-fluid-hiscores" id="${minigame_path}">`;
}