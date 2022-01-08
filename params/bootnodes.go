// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://9f5c18fb1ed3f0f65aed9f470b26044cb686bcf83a896058a41cf43ab08f9bbe6fb9622a94e48c4c6963a95d04505019d773d8a314eb973cd24476e306f44384@49.0.203.177:32668",
	"enode://a87190b5ca90c39a14f65b5318e13a2be7c9d35f8a516547c9e716472c48ce6692870fe1f66ade9f1babfa2d5587d7ffbfa8e5e087e36f8c8a45874b9334ffb9@159.138.92.241:32668",
	"enode://2cf656ebc580ba5fec02c4da908e0f447d4db6fa4d684f3fa59adc59f863c6865b6f9a4657434547c882a13dd41d29b5c9b15d68fe475762a45236b4d31f12c0@114.119.189.247:32668",
	"enode://bc7373b24eaf84f0196a596474d8e92c62713027a7451fe2a05000c50615d078707477765a9e90c946214e53b3d9ad2e3e30f2968bdd591272dc761909920419@202.165.25.187:32668",
	"enode://3653883689d242d379b915877a64c73138ca81a8467c60a9a417ef8858c0ee753f681270045633ed034e50b0d0a367dbb7eaeaf9f2ba964800e03b607d24be7f@202.165.25.154:32668",
	"enode://79032cc19173a2fde8e99604ad7d5099e08e086f1885370cf6d8057de34dba10e2530019fb4183210ae75e03badceea2b1b0ffbfd0481b4d3e5c2556923e03e6@202.165.24.192:32668",

	// "enode://22bc9adaed8405783adc5bb66d4c4349008474b29fac618ad840f35224b7c856d39a3c37cd847acf41718d96832e6c38afdc3bedfc19ade08941850d7f89decb@202.165.24.192:32668",
	// "enode://c7490cf77404fcd6834d5de05e4f9db8c47470dba33c2a3f837e0902e579378fa8609b3ec2ba72be728240207ad7cb9c4ad0fe882b2135178f4e475827989ae8@202.165.25.154:32668",
	// "enode://164f57d3c3d2a75c4d4a80b7cde3508294d6510d2552f10340142ebcb05f02f6806c0eb8af577b6fb74fff49e19843468dbea8cf12030858ab941ed9d55ad7d7@202.165.25.154:32668",

	// "enode://b4bde94fc6295e6b1f6c75a53d5a7ffcf9584364c70d43834e639e1d28c9944b8f44c238f3cf7c6e1e92bbbf07356989f76104a6c759d1b77a8e914213b707f1@34.64.75.216:32668",
	// "enode://8cda0752e7d736f8d0fa0e77ae24d2b84797867f895ae48c8e23195ea16ccb28e00f3a95b8f627b25b7ba2ca5f5a4b769b2390e3043a1e868053e196610a491c@34.92.93.119:32668",
	// "enode://2c97b2b53f9cfe70b575f633fbd7c0bffb3210a4fdd1cf5f29356c5bfb79418961b2afa931c1b88427577ca68bb480fb565a1845b0c95bc8dd3fb4e742e93063@47.108.159.32:32668",
	// "enode://69b036a55eecdbaff7338415afd57b7ea36b7dad5ee7088d70061807e5869db325660d18c289b896e71112cbdf7649b449dd00308a93f1d58cb1d9e62a68ea13@47.108.138.9:32668",
	// Ethereum Foundation Go Bootnodes
	// "enode://7bed18c87054f807bc9096501bc78f737363f357af831791bab07c4fa6c5a1a67cdcf0a097dc2cc918262ef04fb1c05c26026df5c11a6a56666f9b1fb4072210@18.178.30.66:32668",
	// "enode://d67251dd3b050e555679a8abdc427a4c78a9bae174f2fd3b9163c364d27b6a69688ee067cd3214e8ceb71e6e602fd812797b085ae37ed3bf93b78e2b77ae3306@18.181.40.7:32668",
	// "enode://f88bb1f5d0e42cf75ec879212b7c8477d605315d5296fba02bc4600eccf73c64427de46567a320d00985d5bc612168817ba6dff169bd6a4774e112e6db0ff6a2@18.176.66.118:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	// "enode://924543a43d18bc5759a8bdcd17fa9c7c35df63968e9333640b80b58dab94b17a012371c9d46bed10ce7508a607cac76828ca04685893958eee44ade83b856dc2@47.242.237.63:32668",
	// "enode://ebad898d980b520ef6adb54ffb6a68117686e7332f1ea01f7551b7a296a34dd945445a078d7cad019d864c5ef0e0b7f2b5777d94f93adf7dc59f798af72609ac@47.242.235.121:32668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
