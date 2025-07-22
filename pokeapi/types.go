package pokeapi

type NamedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UnnamedURL struct {
	URL string `json:"url"`
}

type Berry struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	GrowthTime       int              `json:"growth_time"`
	MaxHarvest       int              `json:"max_harvest"`
	NaturalGiftPower int              `json:"natural_gift_power"`
	Size             int              `json:"size"`
	Smoothness       int              `json:"smoothness"`
	SoilDryness      int              `json:"soil_dryness"`
	Firmness         NamedURL         `json:"firmness"`
	Flavors          []berryFlavorMap `json:"flavors"`
	Item             NamedURL         `json:"item"`
	NaturalGiftType  NamedURL         `json:"natural_gift_type"`
}

type BerryFirmness struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Berries []NamedURL `json:"berries"`
	Names   []Name     `json:"names"`
}

type berryFlavorMap struct {
	Potency int      `json:"potency"`
	Flavor  NamedURL `json:"flavor"`
}

type BerryFlavor struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Berries     []flavorBerryMap `json:"berries"`
	ContestType NamedURL         `json:"contest_type"`
	Names       []Name           `json:"names"`
}

type BerryFlavorURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type flavorBerryMap struct {
	Potency int      `json:"potency"`
	Berry   NamedURL `json:"berry"`
}

type ContestType struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	BerryFlavor NamedURL      `json:"berry_flavor"`
	Names       []ContestName `json:"names"`
}

type ContestName struct {
	Name     string   `json:"name"`
	Color    string   `json:"color"`
	Language NamedURL `json:"language"`
}

type ContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	Moves             []NamedURL   `json:"moves"`
}

type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []Name `json:"names"`
}

type EncounterCondition struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Values []NamedURL `json:"values"`
	Names  []Name     `json:"names"`
}

type EncounterConditionValue struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Condition NamedURL `json:"condition"`
	Names     []Name   `json:"names"`
}

type EvolutionChain struct {
	ID              int       `json:"id"`
	BabyTriggerItem NamedURL  `json:"baby_trigger_item"`
	Chain           ChainLink `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool              `json:"is_baby"`
	Species          NamedURL          `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item                  NamedURL `json:"item"`
	Trigger               NamedURL `json:"trigger"`
	Gender                int      `json:"gender"`
	HeldItem              NamedURL `json:"held_item"`
	KnownMove             NamedURL `json:"known_move"`
	KnownMoveType         NamedURL `json:"known_move_type"`
	Location              NamedURL `json:"location"`
	MinLevel              int      `json:"min_level"`
	MinHappiness          int      `json:"min_happiness"`
	MinBeauty             int      `json:"min_beauty"`
	MinAffection          int      `json:"min_affection"`
	NeedsOverworldRain    bool     `json:"needs_overworld_rain"`
	PartySpecies          NamedURL `json:"party_species"`
	PartyType             NamedURL `json:"party_type"`
	RelativePhysicalStats int      `json:"relative_physical_stats"`
	TimeOfDay             string   `json:"time_of_day"`
	TradeSpecies          NamedURL `json:"trade_species"`
	TurnUpsideDown        bool     `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Names          []Name     `json:"names"`
	PokemonSpecies []NamedURL `json:"pokemon_species"`
}

type Generation struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Abilities      []NamedURL `json:"abilities"`
	Names          []Name     `json:"names"`
	MainRegion     NamedURL   `json:"main_region"`
	Moves          []NamedURL `json:"moves"`
	PokemonSpecies []NamedURL `json:"pokemon_species"`
	Types          []NamedURL `json:"types"`
	VersionGroups  []NamedURL `json:"version_groups"`
}

type Pokedex struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	IsMainSeries   bool           `json:"is_main_series"`
	Descriptions   []Description  `json:"descriptions"`
	Names          []Name         `json:"names"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	Region         NamedURL       `json:"region"`
	VersionGroups  []NamedURL     `json:"version_groups"`
}

type PokemonEntry struct {
	EntryNumber    int      `json:"entry_number"`
	PokemonSpecies NamedURL `json:"pokemon_species"`
}

type Version struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Names        []Name   `json:"names"`
	VersionGroup NamedURL `json:"version_group"`
}

type VersionGroup struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	Order            int        `json:"order"`
	Generation       NamedURL   `json:"generation"`
	MoveLearnMethods []NamedURL `json:"move_learn_methods"`
	Pokedexes        []NamedURL `json:"pokedexes"`
	Regions          []NamedURL `json:"regions"`
	Versions         []NamedURL `json:"versions"`
}

type Item struct {
	ID                int                      `json:"id"`
	Name              string                   `json:"name"`
	Cost              int                      `json:"cost"`
	FlingPower        int                      `json:"fling_power"`
	FlingEffect       NamedURL                 `json:"fling_effect"`
	Attributes        []NamedURL               `json:"attributes"`
	Category          NamedURL                 `json:"category"`
	EffectEntries     []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices       []GenerationGameIndex    `json:"game_indices"`
	Names             []Name                   `json:"names"`
	Sprites           ItemSprites              `json:"sprites"`
	HeldByPokemon     []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerFor    UnnamedURL               `json:"baby_trigger_for"`
	Machines          []MachineVersionDetail   `json:"machines"`
}

type ItemSprites struct {
	Default string `json:"default"`
}

type ItemHolderPokemon struct {
	Pokemon       NamedURL                         `json:"pokemon"`
	VersionDetail []ItemHolderPokemonVersionDetail `json:"version_details"`
}

type ItemHolderPokemonVersionDetail struct {
	Version NamedURL `json:"version"`
	Rarity  int      `json:"rarity"`
}

type ItemAttribute struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Items        []NamedURL    `json:"items"`
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"descriptions"`
}

type ItemCategory struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Items  []NamedURL `json:"items"`
	Names  []Name     `json:"names"`
	Pocket NamedURL   `json:"pocket"`
}

type ItemFlingEffect struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	EffectEntries []Effect   `json:"effect_entries"`
	Items         []NamedURL `json:"items"`
}

type ItemPocket struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Categories []NamedURL `json:"categories"`
	Names      []Name     `json:"names"`
}

type Location struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Region      NamedURL              `json:"region"`
	Names       []Name                `json:"names"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedURL            `json:"areas"`
}

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedURL              `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedURL                  `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int      `json:"rate"`
	Version NamedURL `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedURL                 `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore      int      `json:"base_score"`
	Rate           int      `json:"rate"`
	PokemonSpecies NamedURL `json:"pokemon_species"`
}

type Region struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Locations      []NamedURL `json:"locations"`
	MainGeneration NamedURL   `json:"main_generation"`
	Names          []Name     `json:"names"`
	Pokedexes      []NamedURL `json:"pokedexes"`
	VersionGroups  []NamedURL `json:"version_groups"`
}

type Machine struct {
	ID           int      `json:"id"`
	Item         NamedURL `json:"item"`
	VersionGroup NamedURL `json:"version_group"`
	Move         NamedURL `json:"move"`
}

type Move struct {
	ID                 int                    `json:"id"`
	Name               string                 `json:"name"`
	Accuracy           int                    `json:"accuracy"`
	EffectChance       int                    `json:"effect_chance"`
	PP                 int                    `json:"pp"`
	Priority           int                    `json:"priority"`
	Power              int                    `json:"power"`
	ContestCombos      ContestComboSet        `json:"contest_combos"`
	ContestType        NamedURL               `json:"contest_type"`
	ContestEffect      UnnamedURL             `json:"contest_effect"`
	DamageClass        NamedURL               `json:"damage_class"`
	EffectEntries      []VerboseEffect        `json:"effect_entries"`
	EffectChanges      []AbilityEffectChange  `json:"effect_changes"`
	LearnedByPokemon   []NamedURL             `json:"learned_by_pokemon"`
	FlavorTextEntries  []MoveFlavorText       `json:"flavor_text_entries"`
	Generation         NamedURL               `json:"generation"`
	Machines           []MachineVersionDetail `json:"machines"`
	Meta               MoveMeta               `json:"meta"`
	Names              []Name                 `json:"names"`
	PastValues         []MovePastValue        `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect UnnamedURL             `json:"super_contest_effect"`
	Target             NamedURL               `json:"target"`
	Type               NamedURL               `json:"type"`
}

type ContestComboSet struct {
	Normal ContestComboDetails `json:"normal"`
	Super  ContestComboDetails `json:"super"`
}

type ContestComboDetails struct {
	UseBefore []NamedURL `json:"use_before"`
	UseAfter  []NamedURL `json:"use_after"`
}

type MoveFlavorText struct {
	FlavorText   string   `json:"flavor_text"`
	Language     NamedURL `json:"language"`
	VersionGroup NamedURL `json:"version_group"`
}

type MoveMeta struct {
	Ailment       NamedURL `json:"ailment"`
	Category      NamedURL `json:"category"`
	MinHits       int      `json:"min_hits"`
	MaxHits       int      `json:"max_hits"`
	MinTurns      int      `json:"min_turns"`
	MaxTurns      int      `json:"max_turns"`
	Drain         int      `json:"drain"`
	Healing       int      `json:"healing"`
	CritRate      int      `json:"crit_rate"`
	AilmentChance int      `json:"ailment_chance"`
	FlinchChance  int      `json:"flinch_chance"`
	StatChance    int      `json:"stat_chance"`
}

type MoveStatChange struct {
	Change int      `json:"change"`
	Stat   NamedURL `json:"stat"`
}

type MovePastValue struct {
	Accuracy      int             `json:"accuracy"`
	Power         int             `json:"power"`
	PP            int             `json:"pp"`
	EffectChance  int             `json:"effect_chance"`
	EffectEntries []VerboseEffect `json:"effect_entries"`
	Type          NamedURL        `json:"type"`
	VersionGroup  NamedURL        `json:"version_group"`
}

type MoveAilment struct {
	ID    int        `json:"id"`
	Name  string     `json:"name"`
	Moves []NamedURL `json:"moves"`
	Names []Name     `json:"names"`
}

type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

type MoveCategory struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Moves        []NamedURL    `json:"moves"`
	Descriptions []Description `json:"descriptions"`
}

type MoveDamageClass struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Moves        []NamedURL    `json:"moves"`
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"descriptions"`
}

type MoveLearnMethod struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Names         []Name        `json:"names"`
	Descriptions  []Description `json:"descriptions"`
	VersionGroups []NamedURL    `json:"version_groups"`
}

type MoveTarget struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Moves        []NamedURL    `json:"moves"`
	Descriptions []Description `json:"descriptions"`
	Names        []Name        `json:"names"`
}

type Ability struct {
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	IsMainSeries      bool                  `json:"is_main_series"`
	Generation        NamedURL              `json:"generation"`
	Names             []Name                `json:"names"`
	EffectEntries     []VerboseEffect       `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

type AbilityEffectChange struct {
	EffectEntries []Effect `json:"effect_entries"`
	VersionGroup  NamedURL `json:"version_group"`
}

type AbilityFlavorText struct {
	FlavorText   string   `json:"flavor_text"`
	Language     NamedURL `json:"language"`
	VersionGroup NamedURL `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Pokemon  NamedURL `json:"pokemon"`
}

type Characteristic struct {
	ID             int           `json:"id"`
	GeneMod        int           `json:"gene_modulo"`
	PossibleValues []int         `json:"possible_values"`
	HighestStat    NamedURL      `json:"highest_stat"`
	Descriptions   []Description `json:"descriptions"`
}

type EggGroup struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Names          []Name     `json:"names"`
	PokemonSpecies []NamedURL `json:"pokemon_species"`
}

type Gender struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	PokemonSpeciesDetail []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution []NamedURL             `json:"required_for_evolution"`
}

type PokemonSpeciesGender struct {
	Rate           int      `json:"rate"`
	PokemonSpecies NamedURL `json:"pokemon_species"`
}

type GrowthRate struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Formula        string            `json:"formula"`
	Descriptions   []Description     `json:"descriptions"`
	Levels         []GrowthRateLevel `json:"levels"`
	PokemonSpecies []NamedURL        `json:"pokemon_species"`
}

type GrowthRateLevel struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`
}

type Nature struct {
	ID                         int                           `json:"id"`
	Name                       string                        `json:"name"`
	DecreasedStat              NamedURL                      `json:"decreased_stat"`
	IncreasedStat              NamedURL                      `json:"increased_stat"`
	HatesFlavor                NamedURL                      `json:"hates_flavor"`
	LikesFlavor                NamedURL                      `json:"likes_flavor"`
	PokemonStatChanges         []NatureStatChange            `json:"pokemon_nature_stat_changes"`
	MoveBattleStylePreferences []NatureBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []Name                        `json:"names"`
}

type NatureStatChange struct {
	MaxChange      int      `json:"max_change"`
	PokeathlonStat NamedURL `json:"pokeathlon_stat"`
}

type NatureBattleStylePreference struct {
	LowHPPreference  int      `json:"low_hp_preference"`
	HighHPPreference int      `json:"high_hp_preference"`
	MoveBattleStyle  NamedURL `json:"move_battle_style"`
}

type PokeathlonStat struct {
	ID               int                           `json:"id"`
	Name             string                        `json:"name"`
	Names            []Name                        `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSet `json:"affecting_natures"`
}

type NaturePokeathlonStatAffectSet struct {
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

type NaturePokeathlonStatAffect struct {
	Nature    NamedURL `json:"nature"`
	MaxChange int      `json:"max_change"`
}

type Pokemon struct {
	ID                     int                  `json:"id"`
	Name                   string               `json:"name"`
	BaseExperience         int                  `json:"base_experience"`
	Height                 int                  `json:"height"`
	IsDefault              bool                 `json:"is_default"`
	Order                  int                  `json:"order"`
	Weight                 int                  `json:"weight"`
	Abilities              []PokemonAbility     `json:"abilities"`
	Forms                  []NamedURL           `json:"forms"`
	GameIndices            []VersionGameIndex   `json:"game_indices"`
	HeldItems              []PokemonHeldItem    `json:"held_items"`
	LocationAreaEncounters string               `json:"location_area_encounters"`
	Moves                  []PokemonMove        `json:"moves"`
	PastTypes              []PokemonTypePast    `json:"past_types"`
	PastAbilities          []PokemonAbilityPast `json:"past_abilities"`
	Sprites                PokemonSprites       `json:"sprites"`
	Cries                  PokemonCries         `json:"cries"`
	Species                NamedURL             `json:"species"`
	Stats                  []PokemonStat        `json:"stats"`
	Types                  []PokemonType        `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  NamedURL `json:"ability"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type NamedURL `json:"type"`
}

type PokemonFormType struct {
	Slot int      `json:"slot"`
	Type NamedURL `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedURL      `json:"generation"`
	Types      []PokemonType `json:"types"`
}

type PokemonAbilityPast struct {
	Generation NamedURL         `json:"generation"`
	Abilities  []PokemonAbility `json:"abilities"`
}

type PokemonHeldItem struct {
	Item           NamedURL                       `json:"item"`
	VersionDetails []PokemonHeldItemVersionDetail `json:"version_details"`
}

type PokemonHeldItemVersionDetail struct {
	Rarity  int      `json:"rarity"`
	Version NamedURL `json:"version"`
}

type PokemonMove struct {
	Move                NamedURL                        `json:"move"`
	VersionGroupDetails []PokemonMoveVersionGroupDetail `json:"version_group_details"`
}

type PokemonMoveVersionGroupDetail struct {
	MoveLearnMethod NamedURL `json:"move_learn_method"`
	VersionGroup    NamedURL `json:"version_group"`
	LevelLearnedAt  int      `json:"level_learned_at"`
	Order           int      `json:"order"`
}

type PokemonStat struct {
	Stat     NamedURL `json:"stat"`
	Effort   int      `json:"effort"`
	BaseStat int      `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonLocationArea struct {
	LocationArea   NamedURL                 `json:"location_area"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PokemonColor struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Names          []Name     `json:"names"`
	PokemonSpecies []NamedURL `json:"pokemon_species"`
}

type PokemonForm struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Order        int                `json:"order"`
	FormOrder    int                `json:"form_order"`
	IsDefault    bool               `json:"is_default"`
	IsBattleOnly bool               `json:"is_battle_only"`
	IsMega       bool               `json:"is_mega"`
	FormName     string             `json:"form_name"`
	Pokemon      NamedURL           `json:"pokemon"`
	Types        []PokemonFormType  `json:"types"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup NamedURL           `json:"version_group"`
	Names        []Name             `json:"names"`
	FormNames    []Name             `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

type PokemonHabitat struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Names          []Name     `json:"names"`
	PokemonSpecies []NamedURL `json:"pokemon_species"`
}

type PokemonShape struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	AwesomeNames   []AwesomeName `json:"awesome_names"`
	Names          []Name        `json:"names"`
	PokemonSpecies []NamedURL    `json:"pokemon_species"`
}

type AwesomeName struct {
	AwesomeName string   `json:"awesome_name"`
	Language    NamedURL `json:"language"`
}

type PokemonSpecies struct {
	ID                   int                      `json:"id"`
	Name                 string                   `json:"name"`
	Order                int                      `json:"order"`
	GenderRate           int                      `json:"gender_rate"`
	CaptureRate          int                      `json:"capture_rate"`
	BaseHappiness        int                      `json:"base_happiness"`
	IsBaby               bool                     `json:"is_baby"`
	IsLegendary          bool                     `json:"is_legendary"`
	IsMythical           bool                     `json:"is_mythical"`
	HatchCounter         int                      `json:"hatch_counter"`
	HasGenderDifferences bool                     `json:"has_gender_differences"`
	FormsSwitchable      bool                     `json:"forms_switchable"`
	GrowthRate           NamedURL                 `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups            []NamedURL               `json:"egg_groups"`
	Color                NamedURL                 `json:"color"`
	Shape                NamedURL                 `json:"shape"`
	EvolvesFromSpecies   NamedURL                 `json:"evolves_from_species"`
	EvolutionChain       UnnamedURL               `json:"evolution_chain"`
	Habitat              NamedURL                 `json:"habitat"`
	Generation           NamedURL                 `json:"generation"`
	Names                []Name                   `json:"names"`
	PalParkEncounters    []PalParkEncounterArea   `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText             `json:"flavor_text_entries"`
	FormDescriptions     []Description            `json:"form_descriptions"`
	Genera               []Genus                  `json:"genera"`
	Varieties            []PokemonSpeciesVariety  `json:"varieties"`
}

type Genus struct {
	Genus    string   `json:"genus"`
	Language NamedURL `json:"language"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int      `json:"entry_number"`
	Pokedex     NamedURL `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore int      `json:"base_score"`
	Rate      int      `json:"rate"`
	Area      NamedURL `json:"area"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool     `json:"is_default"`
	Pokemon   NamedURL `json:"pokemon"`
}

type Stat struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	GameIndex        int                  `json:"game_index"`
	IsBattleOnly     bool                 `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics  []NamedURL           `json:"characteristics"`
	MoveDamageClass  NamedURL             `json:"move_damage_class"`
	Names            []Name               `json:"names"`
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}

type MoveStatAffect struct {
	Move   NamedURL `json:"move"`
	Change int      `json:"change"`
}

type NatureStatAffectSets struct {
	Increase []NamedURL `json:"increase"`
	Decrease []NamedURL `json:"decrease"`
}

type Type struct {
	ID                  int                      `json:"id"`
	Name                string                   `json:"name"`
	DamageRelations     TypeRelations            `json:"damage_relations"`
	PastDamageRelations []TypePastDamageRelation `json:"past_damage_relations"`
	GameIndices         []GenerationGameIndex    `json:"game_indices"`
	Generation          NamedURL                 `json:"generation"`
	MoveDamageClass     NamedURL                 `json:"move_damage_class"`
	Names               []Name                   `json:"names"`
	Pokemon             []TypePokemon            `json:"pokemon"`
	Moves               []NamedURL               `json:"moves"`
}

type TypePokemon struct {
	Slot    int      `json:"slot"`
	Pokemon NamedURL `json:"pokemon"`
}

type TypeRelations struct {
	NoDamageTo       []NamedURL `json:"no_damage_to"`
	HalfDamageTo     []NamedURL `json:"half_damage_to"`
	DoubleDamageTo   []NamedURL `json:"double_damage_to"`
	NoDamageFrom     []NamedURL `json:"no_damage_from"`
	HalfDamageFrom   []NamedURL `json:"half_damage_from"`
	DoubleDamageFrom []NamedURL `json:"double_damage_from"`
}

type TypePastDamageRelation struct {
	Generation      NamedURL      `json:"generation"`
	DamageRelations TypeRelations `json:"damage_relations"`
}

type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Iso639   string `json:"iso639"`
	Iso3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

type Description struct {
	Description string   `json:"description"`
	Language    NamedURL `json:"language"`
}

type Effect struct {
	Effect   string   `json:"effect"`
	Language NamedURL `json:"language"`
}

type Encounter struct {
	MinLevel        int        `json:"min_level"`
	MaxLevel        int        `json:"max_level"`
	ConditionValues []NamedURL `json:"condition_values"`
	Chance          int        `json:"chance"`
	Method          NamedURL   `json:"method"`
}

type FlavorText struct {
	FlavorText string   `json:"flavor_text"`
	Language   NamedURL `json:"language"`
	Version    NamedURL `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int      `json:"game_index"`
	Generation NamedURL `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      NamedURL `json:"machine"`
	VersionGroup NamedURL `json:"version_group"`
}

type Name struct {
	Name     string   `json:"name"`
	Language NamedURL `json:"language"`
}

type VerboseEffect struct {
	Effect      string   `json:"effect"`
	ShortEffect string   `json:"short_effect"`
	Language    NamedURL `json:"language"`
}

type VersionEncounterDetail struct {
	Version          NamedURL    `json:"version"`
	MaxChance        int         `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type VersionGameIndex struct {
	GameIndex int      `json:"game_index"`
	Version   NamedURL `json:"version"`
}

type VersionGroupFlavorText struct {
	Text         string   `json:"text"`
	Language     NamedURL `json:"language"`
	VersionGroup NamedURL `json:"version_group"`
}
