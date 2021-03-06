package protocol_impl

// Constant packet IDs to avoid magic numbers
// Packet + [State] + [Binding (Cl = Clientbound; Sv = Serverbound)] + [Name]
const (
	PacketHandshakeSvHandshake            = 0x00
	PacketHandshakeSvLegacyServerListPing = 0xFE

	PacketStatusClResponse = 0x00
	PacketStatusClPong     = 0x01
	PacketStatusSvRequest  = 0x00
	PacketStatusSvPing     = 0x01

	PacketLoginClDisconnect          = 0x00
	PacketLoginClEncryptionRequest   = 0x01
	PacketLoginClLoginSuccess        = 0x02
	PacketLoginClSetCompression      = 0x03
	PacketLoginClLoginPluginRequest  = 0x04
	PacketLoginSvLoginStart          = 0x00
	PacketLoginSvEncryptionResponse  = 0x01
	PacketLoginSvLoginPluginResponse = 0x02

	PacketPlayClSpawnEntity                = 0x00
	PacketPlayClSpawnExperienceOrb         = 0x01
	PacketPlayClSpawnWeatherEntity         = 0x02
	PacketPlayClSpawnLivingEntity          = 0x03
	PacketPlayClSpawnPainting              = 0x04
	PacketPlayClSpawnPlayer                = 0x05
	PacketPlayClEntityAnimation            = 0x06
	PacketPlayClStatistics                 = 0x07
	PacketPlayClAcknowledgePlayerDigging   = 0x08
	PacketPlayClBlockBreakAnimation        = 0x09
	PacketPlayClBlockEntityData            = 0x0A
	PacketPlayClBlockAction                = 0x0B
	PacketPlayClBlockChange                = 0x0C
	PacketPlayClBossBar                    = 0x0D
	PacketPlayClServerDifficulty           = 0x0E
	PacketPlayClChatMessage                = 0x0F
	PacketPlayClMultiBlockChange           = 0x10
	PacketPlayClTabComplete                = 0x11
	PacketPlayClDeclareCommands            = 0x12
	PacketPlayClWindowConfirmation         = 0x13
	PacketPlayClCloseWindow                = 0x14
	PacketPlayClWindowItems                = 0x15
	PacketPlayClWindowProperty             = 0x16
	PacketPlayClSetSlot                    = 0x17
	PacketPlayClSetCooldown                = 0x18
	PacketPlayClPluginMessage              = 0x19
	PacketPlayClNamedSoundEffect           = 0x1A
	PacketPlayClDisconnect                 = 0x1B
	PacketPlayClEntityStatus               = 0x1C
	PacketPlayClExplosions                 = 0x1D
	PacketPlayClUnloadChunk                = 0x1E
	PacketPlayClChangeGameState            = 0x1F
	PacketPlayClOpenHorseWindow            = 0x20
	PacketPlayClKeepAlive                  = 0x21
	PacketPlayClChunkData                  = 0x22
	PacketPlayClEffect                     = 0x23
	PacketPlayClParticle                   = 0x24
	PacketPlayClUpdateLight                = 0x25
	PacketPlayClJoinGame                   = 0x26
	PacketPlayClMapData                    = 0x27
	PacketPlayClTradeList                  = 0x28
	PacketPlayClEntityPosition             = 0x29
	PacketPlayClEntityPositionAndRotation  = 0x2A
	PacketPlayClEntityRotation             = 0x2B
	PacketPlayClEntityMovement             = 0x2C
	PacketPlayClVehicleMove                = 0x2D
	PacketPlayClOpenBook                   = 0x2E
	PacketPlayClOpenWindow                 = 0x2F
	PacketPlayClOpenSignEditor             = 0x30
	PacketPlayClCraftRecipeResponse        = 0x31
	PacketPlayClPlayerAbilities            = 0x32
	PacketPlayClCombatEvent                = 0x33
	PacketPlayClPlayerInfo                 = 0x34
	PacketPlayClFacePlayer                 = 0x35
	PacketPlayClPlayerPositionAndLook      = 0x36
	PacketPlayClUnlockRecipes              = 0x37
	PacketPlayClDestroyEntities            = 0x38
	PacketPlayClRemoveEntityEffect         = 0x39
	PacketPlayClResourcePackSend           = 0x3A
	PacketPlayClRespawn                    = 0x3B
	PacketPlayClEntityHeadLook             = 0x3C
	PacketPlayClSelectAdvancementTab       = 0x3D
	PacketPlayClWorldBorder                = 0x3E
	PacketPlayClCamera                     = 0x3F
	PacketPlayClHeldItemChange             = 0x40
	PacketPlayClUpdateViewPosition         = 0x41
	PacketPlayClUpdateViewDistance         = 0x42
	PacketPlayClDisplayScoreboard          = 0x43
	PacketPlayClEntityMetadata             = 0x44
	PacketPlayClAttachEntity               = 0x45
	PacketPlayClEntityVelocity             = 0x46
	PacketPlayClEntityEquipment            = 0x47
	PacketPlayClSetExperience              = 0x48
	PacketPlayClUpdateHealth               = 0x49
	PacketPlayClScoreboardObjective        = 0x4A
	PacketPlayClSetPassengers              = 0x4B
	PacketPlayClTeams                      = 0x4C
	PacketPlayClUpdateScore                = 0x4D
	PacketPlayClSpawnPosition              = 0x4E
	PacketPlayClTimeUpdate                 = 0x4F
	PacketPlayClTitle                      = 0x50
	PacketPlayClEntitySoundEffect          = 0x51
	PacketPlayClSoundEffect                = 0x52
	PacketPlayClStopSound                  = 0x53
	PacketPlayClPlayerListHeaderAndFooter  = 0x54
	PacketPlayClNBTQueryResponse           = 0x55
	PacketPlayClCollectItem                = 0x56
	PacketPlayClEntityTeleport             = 0x57
	PacketPlayClAdvancements               = 0x58
	PacketPlayClEntityProperties           = 0x59
	PacketPlayClEntityEffect               = 0x5A
	PacketPlayClDeclareRecipes             = 0x5B
	PacketPlayClTags                       = 0x5C
	PacketPlaySvTeleportConfirm            = 0x00
	PacketPlaySvQueryBlockNBT              = 0x01
	PacketPlaySvSetDifficulty              = 0x02
	PacketPlaySvChatMessage                = 0x03
	PacketPlaySvClientStatus               = 0x04
	PacketPlaySvClientSettings             = 0x05
	PacketPlaySvTabComplete                = 0x06
	PacketPlaySvWindowConfirmation         = 0x07
	PacketPlaySvClickWindowButton          = 0x08
	PacketPlaySvClickWindow                = 0x09
	PacketPlaySvCloseWindow                = 0x0A
	PacketPlaySvPluginMessage              = 0x0B
	PacketPlaySvEditBook                   = 0x0C
	PacketPlaySvEntityNBTRequest           = 0x0D
	PacketPlaySvInteractEntity             = 0x0E
	PacketPlaySvKeepAlive                  = 0x0F
	PacketPlaySvLockDifficulty             = 0x10
	PacketPlaySvPlayerPosition             = 0x11
	PacketPlaySvPlayerPositionAndRotation  = 0x12
	PacketPlaySvPlayerRotation             = 0x13
	PacketPlaySvPlayerMovement             = 0x14
	PacketPlaySvVehicleMove                = 0x15
	PacketPlaySvSteerBoat                  = 0x16
	PacketPlaySvPickItem                   = 0x17
	PacketPlaySvCraftRecipeRequest         = 0x18
	PacketPlaySvPlayerAbilities            = 0x19
	PacketPlaySvPlayerDigging              = 0x1A
	PacketPlaySvEntityAction               = 0x1B
	PacketPlaySvSteerVehicle               = 0x1C
	PacketPlaySvRecipeBookData             = 0x1D
	PacketPlaySvNameItem                   = 0x1E
	PacketPlaySvResourcePackStatus         = 0x1F
	PacketPlaySvAdvancementTab             = 0x20
	PacketPlaySvSelectTrade                = 0x21
	PacketPlaySvSetBeaconEffect            = 0x22
	PacketPlaySvHeldItemChange             = 0x23
	PacketPlaySvUpdateCommandBlock         = 0x24
	PacketPlaySvUpdateCommandBlockMinecart = 0x25
	PacketPlaySvCreativeInventoryAction    = 0x26
	PacketPlaySvUpdateJigsawBlock          = 0x27
	PacketPlaySvUpdateStructureBlock       = 0x28
	PacketPlaySvUpdateSign                 = 0x29
	PacketPlaySvAnimation                  = 0x2A
	PacketPlaySvSpectate                   = 0x2B
	PacketPlaySvPlayerBlockPlacement       = 0x2C
	PacketPlaySvUseItem                    = 0x2D
)
