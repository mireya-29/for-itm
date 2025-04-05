func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]interface{}) string {

	contains := func(slice []string, item string) bool {
		for _, v := range slice {
			if v == item {
				return true
			}
		}
		return false
	}

	if fromMemberCheck, exists := socialGraph[fromMember]; exists {

		fromFollowing := strings.Split(fromMemberCheck["following"].(string), ",")

		if contains(fromFollowing, toMember) {

			if toMemberCheck, exists := socialGraph[toMember]; exists {
				toFollowing := strings.Split(toMemberCheck["following"].(string), ",")

				if contains(toFollowing, fromMember) {
					return "friends"
				}
			}

			return "follower"
		}
	}

	if toMemberCheck, exists := socialGraph[toMember]; exists {
		toFollowing := strings.Split(toMemberCheck["following"].(string), ",")
		if contains(toFollowing, fromMember) {
			return "followed by"
		}
	}

	return "no relationship"
}
