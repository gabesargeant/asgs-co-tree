package main

/**
Levels and code maps
*/

var levelCodeMap = map[string]string{
	"MB":    "MB_CODE_2016",
	"SA1":   "SA1_MAINCODE_2016",
	"SA2":   "SA2_MAINCODE_2016",
	"SA3":   "SA3_CODE_2016",
	"SA4":   "SA4_CODE_2016",
	"GCCSA": "GCCSA_CODE_2016",
	"STE":   "STATE_CODE_2016",
	"AUS":   "AUS_CODE_2016",
	"DZN":   "DZN_CODE_2016",
	"LGA":   "LGA_CODE_2016",
	"POA":   "POA_CODE_2016",
	"ADD":   "ADD_CODE_2016",
	"NRMR":  "NRMR_CODE_2016",
	"SSC":   "SSC_CODE_2016",
	"TR":    "TR_CODE_2016",
	"RA":    "RA_CODE_2016",
	"ILOC":  "ILOC_CODE_2016",
	"IARE":  "IARE_CODE_2016",
	"IREG":  "IREG_CODE_2016",
	"UCL":   "UCL_CODE_2016",
	"SOSR":  "SOSR_CODE_2016",
	"SOS":   "SOS_CODE_2016",
	"SUA":   "SUA_CODE_2016",
	"SED":   "SED_CODE_2016",
	"CED":   "CED_CODE_2016",
}

var levelNameMap = map[string]string{
	"MB":    "MB_CATEGORY_NAME_2016",
	"SA1":   "SA1_NAME_2016",
	"SA2":   "SA2_NAME_2016",
	"SA3":   "SA3_NAME_2016",
	"SA4":   "SA4_NAME_2016",
	"GCCSA": "GCCSA_NAME_2016",
	"STE":   "STATE_NAME_2016",
	"AUS":   "AUS_NAME_2016",
	"DZN":   "DZN_NAME_2016",
	"LGA":   "LGA_NAME_2016",
	"POA":   "POA_NAME_2016",
	"ADD":   "ADD_NAME_2016",
	"NRMR":  "NRMR_NAME_2016",
	"SSC":   "SSC_NAME_2016",
	"TR":    "TR_NAME_2016",
	"RA":    "RA_NAME_2016",
	"ILOC":  "ILOC_NAME_2016",
	"IARE":  "IARE_NAME_2016",
	"IREG":  "IREG_NAME_2016",
	"UCL":   "UCL_NAME_2016",
	"SOSR":  "SOSR_NAME_2016",
	"SOS":   "SOS_NAME_2016",
	"SUA":   "SUA_NAME_2016",
	"SED":   "SED_NAME_2016",
	"CED":   "CED_NAME_2016",
}

var asgsRegionArray = []string{
	"MB_CODE_2016",
	"MB_CATEGORY_NAME_2016",
	"SA1_MAINCODE_2016",
	"SA1_NAME_2016",
	"SA2_MAINCODE_2016",
	"SA2_NAME_2016",
	"SA3_CODE_2016",
	"SA3_NAME_2016",
	"SA4_CODE_2016",
	"SA4_NAME_2016",
	"GCCSA_CODE_2016",
	"GCCSA_NAME_2016",
	"STATE_CODE_2016",
	"STATE_NAME_2016",
	"AUS_CODE_2016",
	"AUS_NAME_2016",
	"DZN_CODE_2016",
	"DZN_NAME_2016",
	"LGA_NAME_2015",
	"LGA_CODE_2015",
	"LGA_CODE_2016",
	"LGA_NAME_2016",
	"POA_CODE_2016",
	"POA_NAME_2016",
	"ADD_CODE_2016",
	"ADD_NAME_2016",
	"NRMR_CODE_2016",
	"NRMR_NAME_2016",
	"SSC_CODE_2016",
	"SSC_NAME_2016",
	"TR_CODE_2016",
	"TR_NAME_2016",
	"RA_CODE_2016",
	"RA_NAME_2016",
	"ILOC_CODE_2016",
	"ILOC_NAME_2016",
	"IARE_CODE_2016",
	"IARE_NAME_2016",
	"IREG_CODE_2016",
	"IREG_NAME_2016",
	"UCL_CODE_2016",
	"UCL_NAME_2016",
	"SOSR_CODE_2016",
	"SOSR_NAME_2016",
	"SOS_CODE_2016",
	"SOS_NAME_2016",
	"SUA_CODE_2016",
	"SUA_NAME_2016",
	"SED_CODE_2016",
	"SED_NAME_2016",
	"CED_CODE_2016",
	"CED_NAME_2016",
}

/**
*
* Sequences
*
**/

var asgsParentSeq = map[string][]string{
	"AUS": {},
	"STE": {"AUS"},
	"SA4": {"STE"},
	"SA3": {"SA4"},
	"SA2": {"SA3"},
	"SA1": {"SA2"},
	"MB":  {"SA1"},
}

var raParentSeq = map[string][]string{
	"AUS": {},
	"STE": {"AUS"},
	"RA":  {"STE"},
	"SA1": {"RA"},
	"MB":  {"SA1"},
}

var sosParentSeq = map[string][]string{
	"AUS":  {},
	"STE":  {"AUS"},
	"SOS":  {"STE"},
	"SOSR": {"SOS"},
	"UCL":  {"SOSR"},
	"SA1":  {"UCL"},
	"MB":   {"SA1"},
}

var iregParentSeq = map[string][]string{
	"AUS":  {},
	"STE":  {"AUS"},
	"IREG": {"STE"},
	"IARE": {"IREG"},
	"ILOC": {"IARE"},
	"SA1":  {"ILOC"},
	"MB":   {"SA1"},
}

var poaParentSeq = map[string][]string{
	"AUS": {},
	"STE": {"AUS"},
	"POA": {"STE"},
	"MB":  {"POA"},
}

var lgaParentSeq = map[string][]string{
	"AUS": {},
	"STE": {"AUS"},
	"LGA": {"STE"},
	"MB":  {"LGA"},
}

var sscParentSeq = map[string][]string{
	"AUS": {},
	"STE": {"AUS"},
	"SSC": {"STE"},
	"MB":  {"SSC"},
}

//Map of maps of string arrays :D
var levelSequenceSets = map[string]map[string][]string{
	"asgs": asgsParentSeq,
	//"ra":   raParentSeq,
	//"sos":  sosParentSeq,
	"ireg": iregParentSeq,
	"poa": poaParentSeq,
	"ssc":  sscParentSeq,
	"lga":  lgaParentSeq,
}
