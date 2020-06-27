package global

/*
 * 存储所有服务的索引值
 */

var (
	registryId = "1"
)

func GetRegistryID() string {
	return registryId
}
