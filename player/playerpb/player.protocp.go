package playerpb

import "github.com/adamhicks/ctrlaltdefeat/player"

func PlayerPartInfoFromProto(info *PlayerPart) player.PartInfo {
	return player.PartInfo{
		Player: info.Player,
		Part:   info.Part,
	}
}

func PlayerPartInfoToProto(info player.PartInfo) *PartInfo {
	return &PartInfo{
		Player: info.Player,
		Part:   info.Part,
	}
}

func PlayerRoundInfoFromProto(info *RoundInfo) player.RoundInfo {
	parts := make([]player.PartInfo, 0, len(info.Parts))
	for i, p := range info.Parts {
		parts[i] = PlayerPartInfoFromProto(p)
	}
	return player.RoundInfo{
		Rank:  info.Rank,
		Parts: parts,
	}
}

func PlayerRoundInfoToProto(info player.RoundInfo) *RoundInfo {
	parts := make([]*PartInfo, 0, len(info.Parts))
	for i, p := range info.Parts {
		parts[i] = PlayerPartInfoToProto(p)
	}
	return &RoundInfo{
		Rank:  info.Rank,
		Parts: parts,
	}
}
