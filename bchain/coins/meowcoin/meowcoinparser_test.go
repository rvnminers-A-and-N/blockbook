//go:build unittest

package meowcoin

import (
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress_Mainnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "P2PKH1",
			args:    args{address: "MPyNGZSSZ4rbjkVJRLn3v64pMcktpEYJnU"},
			want:    "76a914b04db9611672485161d38b4ee215e7a8c57d944988ac",
			wantErr: false,
		},
		{
			name:    "P2PKH2",
			args:    args{address: "MQwkVp9SVLtcNHckGw3iavmyDxD99c7zTz"},
			want:    "76a914baf792cbe04d58ed0770c5fe73065322fa4ceb4d88ac",
			wantErr: false,
		},
		{
			name:    "P2SH1",
			args:    args{address: "rLAWjiJv3kY6fXCFUw3R8jXc21wkvhPzSL"},
			want:    "a91498cc6f0a954dfafb081d7e991c501ee0688bda9a87",
			wantErr: false,
		},
		{
			name:    "P2SH2",
			args:    args{address: "rJr8Cd3bEyyZsaEyauxrb4tnA8Jw9XSjLd"},
			want:    "a9148a5a022dfa5c3ba85dc02074df17492f40be05c487",
			wantErr: false,
		},
	}
	parser := NewMeowcoinParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}

var (
	testTx1       bchain.Tx
	testTxPacked1 = "0a20d4d3a093586eae0c3668fd288d9e24955928a894c20b551b38dd18c99b123a7c12e1010200000001c171348ffc8976074fa064e48598a816fce3798afc635fb67d99580e50b8e614000000006a473044022009e07574fa543ad259bd3334eb365c655c96d310c578b64c24d7f77fa7dc591c0220427d8ae6eacd1ca2d1994e9ec49cb322aacdde98e4bdb065e0fce81162fb3aa9012102d46827546548b9b47ae1e9e84fc4e53513e0987eeb1dd41220ba39f67d3bf46affffffff02f8137114000000001976a914587a2afa560ccaeaeb67cb72a0db7e2573a179e488ace0c48110000000001976a914d85e6ab66ab0b2c4cfd40ca3b0a779529da5799288ac0000000018c7e1b3e50528849128329401122014e6b8500e58997db65f63fc8a79e3fc16a89885e464a04f077689fc8f3471c1226a473044022009e07574fa543ad259bd3334eb365c655c96d310c578b64c24d7f77fa7dc591c0220427d8ae6eacd1ca2d1994e9ec49cb322aacdde98e4bdb065e0fce81162fb3aa9012102d46827546548b9b47ae1e9e84fc4e53513e0987eeb1dd41220ba39f67d3bf46a28ffffffff0f3a450a04147113f81a1976a914587a2afa560ccaeaeb67cb72a0db7e2573a179e488ac222252484d31746d64766b6b3776446f69477877554a414d4e4e6d447179775a3574456e3a470a041081c4e010011a1976a914d85e6ab66ab0b2c4cfd40ca3b0a779529da5799288ac2222525631463939623955424272434d38614e4b7567737173444d3869716f4371374d744002"

	testTx2       bchain.Tx
	testTxPacked2 = "0a208e480d5c1bf7f11d1cbe396ab7dc14e01ea4e1aff45de7c055924f61304ad43412f40202000000029e2e14113b2f55726eebaa440edec707fcec3a31ce28fa125afea1e755fb6850010000006a47304402204034c3862f221551cffb2aa809f621f989a75cdb549c789a5ceb3a82c0bcc21c022001b4638f5d73fdd406a4dd9bf99be3dfca4a572b8f40f09b8fd495a7756c0db70121027a32ef45aef2f720ccf585f6fb0b8a7653db89cacc3320e5b385146851aba705fefffffff3b240ae32c542786876fcf23b4b2ab4c34ef077912898ee529756ed4ba35910000000006a47304402204d442645597b13abb85e96e5acd34eff50a4418822fe6a37ed378cdd24574dff02205ae667c56eab63cc45a51063f15b72136fd76e97c46af29bd28e8c4d405aa211012102cde27d7b29331ea3fef909a8d91f6f7753e99a3dd129914be50df26eed73fab3feffffff028447bf38000000001976a9146d7badec5426b880df25a3afc50e476c2423b34b88acb26b556a740000001976a914b3020d0ab85710151fa509d5d9a4e783903d681888ac83080a0018c7e1b3e505208391282884912832960112205068fb55e7a1fe5a12fa28ce313aecfc07c7de0e44aaeb6e72552f3b11142e9e1801226a47304402204034c3862f221551cffb2aa809f621f989a75cdb549c789a5ceb3a82c0bcc21c022001b4638f5d73fdd406a4dd9bf99be3dfca4a572b8f40f09b8fd495a7756c0db70121027a32ef45aef2f720ccf585f6fb0b8a7653db89cacc3320e5b385146851aba70528feffffff0f32940112201059a34bed569752ee98289177f04ec3b42a4b3bf2fc76687842c532ae40b2f3226a47304402204d442645597b13abb85e96e5acd34eff50a4418822fe6a37ed378cdd24574dff02205ae667c56eab63cc45a51063f15b72136fd76e97c46af29bd28e8c4d405aa211012102cde27d7b29331ea3fef909a8d91f6f7753e99a3dd129914be50df26eed73fab328feffffff0f3a450a0438bf47841a1976a9146d7badec5426b880df25a3afc50e476c2423b34b88ac2222524b4735747057776a6874716464546741335168556837516d4b637576426e6842583a480a05746a556bb210011a1976a914b3020d0ab85710151fa509d5d9a4e783903d681888ac222252526268564d624c6675657a485077554d756a546d4446417a76363459396d4a71644002"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "0200000004007804d6323ba87942a0974ca794d0d07bb112ea365fe68c559565f9bc3e758e000000006a4730440220102438390e11c65145929cfc7b7b46b4faac5be92e02e67f204e149fc0dc18240220440efc8e8119994d96cc54117e65ab86e97e1a88a05f8267116f042c88eaa33401210228c82ca1f4b0a51889847649461afbcb9be813f6529155d2b29e12b9da38654dfeffffff505538d2f7e23df8f443b8408d9bef34ca9d4a2639f9b8b07ea7dd63d7101df0030000006a47304402207cb6f0e714853c1155d885dfb262280bd501942b4720c8cdbc77ed210b482fb802206faf64cd5fded1fc6d6423082039d39b19d6ab3ef3371d5fbb7242427db3b986012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864feffffff6a8384964fe3a5e702ed6835a498f596036d300676f13278062e55fffc4b9fc5030000006a4730440220309cd63e140cfe16e75db2f94377dec7d17b44590e441a8547cda8d41d0776c602206abb1c1243c03abbe1899cbe5028b9f36021a9beae1485111797f428f8dbc138012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864feffffff87ecd629c767e04cd8cfc4e68ef1a7ddf5ee6d2bd8cb4a068057fd69c35d2e80010000006b483045022100fea4689c71c2c71a27c03c10d5a16cc0e64de5fbf79a9c5d261f1de07c24eb65022001663b235fcdb02bfbf167d08f049015ba92020de9fe507962f4caac0d74a5d7012102ea91912db24f808a05aaf649ee35df1e7e3bc37461ee555212433ea303e79b63feffffff0200943577000000001976a91446a3097dadaa554d3199b5a5a37762e0b443457488acb1c22000000000001976a914e4b2ed8c899e42cec60abfbc17c998265ce775bc88accc370200",
		Blocktime: 1671173626,
		Time:      1671173626,
		Txid:      "3b9acd076472d51a4981da9fba7a4e8db1f4e0b4be8eec672c9fe6b5d16f2016",
		LockTime:  145356,
		Version:   2,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "4730440220102438390e11c65145929cfc7b7b46b4faac5be92e02e67f204e149fc0dc18240220440efc8e8119994d96cc54117e65ab86e97e1a88a05f8267116f042c88eaa33401210228c82ca1f4b0a51889847649461afbcb9be813f6529155d2b29e12b9da38654d",
				},
				Txid:     "8e753ebcf96595558ce65f36ea12b17bd0d094a74c97a04279a83b32d6047800",
				Vout:     0,
				Sequence: 4294967294,
			},
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "47304402207cb6f0e714853c1155d885dfb262280bd501942b4720c8cdbc77ed210b482fb802206faf64cd5fded1fc6d6423082039d39b19d6ab3ef3371d5fbb7242427db3b986012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864",
				},
				Txid:     "f01d10d763dda77eb0b8f939264a9dca34ef9b8d40b843f4f83de2f7d2385550",
				Vout:     3,
				Sequence: 4294967294,
			},
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "4730440220309cd63e140cfe16e75db2f94377dec7d17b44590e441a8547cda8d41d0776c602206abb1c1243c03abbe1899cbe5028b9f36021a9beae1485111797f428f8dbc138012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864",
				},
				Txid:     "c59f4bfcff552e067832f17606306d0396f598a43568ed02e7a5e34f9684836a",
				Vout:     3,
				Sequence: 4294967294,
			},
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100fea4689c71c2c71a27c03c10d5a16cc0e64de5fbf79a9c5d261f1de07c24eb65022001663b235fcdb02bfbf167d08f049015ba92020de9fe507962f4caac0d74a5d7012102ea91912db24f808a05aaf649ee35df1e7e3bc37461ee555212433ea303e79b63",
				},
				Txid:     "802e5dc369fd5780064acbd82b6deef5dda7f18ee6c4cfd84ce067c729d6ec87",
				Vout:     1,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(2000000000),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a91446a3097dadaa554d3199b5a5a37762e0b443457488ac",
					Addresses: []string{
						"MELeqsaP957CyTBTYNaaXm51EMtg3di5r8",
					},
				},
			},
			{
				ValueSat: *big.NewInt(2146993),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914e4b2ed8c899e42cec60abfbc17c998265ce775bc88ac",
					Addresses: []string{
						"MUkQeJDtGDUcju1XHsB62ksNuCsS3GDw7T",
					},
				},
			},
		},
	}

	testTx2 = bchain.Tx{
		Hex:       "02000000025a9c8a62c9a319bb8fcc0a8191666b98c721033d49fb40d57bf1bcf8c9883e13030000006b483045022100eb056f7914ad3609bc68e1b8e72c2f9a0007f54adf40c467dee3711c020a1f8b0220018be04ac9404fdcc7b0765b0a8bb3610b3caed6fcb758aee37b049bf7e56d74012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864feffffffd3186a72199d886dd60039395e2808794e6d6479964017691d028a86b9b8260f030000006a47304402207fcc9d0084f4451b76b053cce92be8752ba43b5b78a7a60259c271ea672f406102201780dcb8cd14ccb0b2acb98fa8f4e64ba55adfca98805fc16b96e37345ce3fc8012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864feffffff02002f6859000000001976a914c1904fcbdd06b8de303ef34a03c9f3ad3683205688ac3162381b000000001976a914a93c7aea9f4c70079562f0855552868ea3f30d5c88acc5370200",
		Blocktime: 1671175212,
		Time:      1671175212,
		Txid:      "c62381bb9549886d967035046d08c4ea792e4b8a2951a8e6004154bd5c9ac55c",
		LockTime:  145349,
		Version:   2,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100eb056f7914ad3609bc68e1b8e72c2f9a0007f54adf40c467dee3711c020a1f8b0220018be04ac9404fdcc7b0765b0a8bb3610b3caed6fcb758aee37b049bf7e56d74012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864",
				},
				Txid:     "133e88c9f8bcf17bd540fb493d0321c7986b6691810acc8fbb19a3c9628a9c5a",
				Vout:     3,
				Sequence: 4294967294,
			},
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "47304402207fcc9d0084f4451b76b053cce92be8752ba43b5b78a7a60259c271ea672f406102201780dcb8cd14ccb0b2acb98fa8f4e64ba55adfca98805fc16b96e37345ce3fc8012103e5a2d2bcebe1585bc288dfc523b279cb521e01dbff2671230b8636667aa66864",
				},
				Txid:     "0f26b8b9868a021d6917409679646d4e7908285e393900d66d889d19726a18d3",
				Vout:     3,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(1500000000),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914c1904fcbdd06b8de303ef34a03c9f3ad3683205688ac",
					Addresses: []string{
						"MRYdXQXcedvir2J1tHzWd2mDTzpRtMjqwG",
					},
				},
			},
			{
				ValueSat: *big.NewInt(456679985),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914a93c7aea9f4c70079562f0855552868ea3f30d5c88ac",
					Addresses: []string{
						"MPKzt5oACQG32EryHxXk1mMc2Gn7pSUFQa",
					},
				},
			},
		},
	}
}

func Test_PackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *MeowcoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "meowcoin-1",
			args: args{
				tx:        testTx1,
				height:    145357,
				blockTime: 1671173626,
				parser:    NewMeowcoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
		{
			name: "meowcoin-2",
			args: args{
				tx:        testTx2,
				height:    145375,
				blockTime: 1671175212,
				parser:    NewMeowcoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_UnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *MeowcoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "meowcoin-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewMeowcoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   145357,
			wantErr: false,
		},
		{
			name: "meowcoin-2",
			args: args{
				packedTx: testTxPacked2,
				parser:   NewMeowcoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx2,
			want1:   145375,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
