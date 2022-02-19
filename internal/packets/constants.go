package packets

var packets = []Packet{
	Motion, Session, LapData, Event, Participants, CarSetups, CarTelemetry,
	CarStatus, FinalClassification, LobbyInfo, CarDamage, SessionHistory,
}

var (
	Motion = Packet{
		Id:          0,
		Name:        "Motion",
		Description: "Contains all motion data for player’s car – only sent while player is in control.",
	}
	Session = Packet{
		Id:          1,
		Name:        "Session",
		Description: "Data about the session – track, time left.",
	}
	LapData = Packet{
		Id:          2,
		Name:        "Lap Data",
		Description: "Data about all the lap times of cars in the session.",
	}
	Event = Packet{
		Id:          3,
		Name:        "Event",
		Description: "Various notable events that happen during a session.",
	}
	Participants = Packet{
		Id:          4,
		Name:        "Participants",
		Description: "List of participants in the session, mostly relevant for multiplayer.",
	}
	CarSetups = Packet{
		Id:          5,
		Name:        "Car Setups",
		Description: "Packet detailing car setups for cars in the race.",
	}
	CarTelemetry = Packet{
		Id:          6,
		Name:        "Car Telemetry",
		Description: "Telemetry data for all cars.",
	}
	CarStatus = Packet{
		Id:          7,
		Name:        "Car Status",
		Description: "Status data for all cars such as damage.",
	}
	FinalClassification = Packet{
		Id:          8,
		Name:        "Final Classification",
		Description: "Final classification confirmation at the end of a race.",
	}
	LobbyInfo = Packet{
		Id:          9,
		Name:        "Lobby Info",
		Description: "Information about players in a multiplayer lobby.",
	}
	CarDamage = Packet{
		Id:          10,
		Name:        "Car Damage",
		Description: "Damage status for all cars.",
	}
	SessionHistory = Packet{
		Id:          11,
		Name:        "Session History",
		Description: "Lap and tyre data for session.",
	}
)
