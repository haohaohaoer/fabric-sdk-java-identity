package org.app.network;

import org.app.config.Config;
import org.hyperledger.fabric.sdk.Peer;

public class PeerInfo {

    Peer[] peers_org1;
    String[] org1_peers_name;
    String[] org1_peers_url;

    public PeerInfo(){
        this.peers_org1 = new Peer[5];
        this.org1_peers_name = new String[]{Config.ORG1_PEER_0,
                Config.ORG1_PEER_1, Config.ORG1_PEER_2, Config.ORG1_PEER_3, Config.ORG1_PEER_4};
        this.org1_peers_url = new String[]{Config.ORG1_PEER_0_URL,
                Config.ORG1_PEER_1_URL, Config.ORG1_PEER_2_URL, Config.ORG1_PEER_3_URL, Config.ORG1_PEER_4_URL};
    }
}
