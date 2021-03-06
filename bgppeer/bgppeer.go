/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package bgppeer

import (
	"context"

	calicoApi "github.com/projectcalico/libcalico-go/lib/apis/v3"
	calicoClient "github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projectcalico/libcalico-go/lib/options"

	"github.com/prometheus/common/log"
)

type BGPPeer struct {
	CalicoClient calicoClient.Interface
}

func (b *BGPPeer) ListBGPPeers() (*calicoApi.BGPPeerList, error) {
	return b.CalicoClient.BGPPeers().List(context.Background(), options.ListOptions{})
}

func (b *BGPPeer) SaveBGPPeer(peer *calicoApi.BGPPeer) error {
	if peer.GetUID() == "" {
		log.Debugf("Creating new BGPPeers: %s", peer.Name)
		if _, err := b.CalicoClient.BGPPeers().Create(context.Background(), peer, options.SetOptions{}); err != nil {
			return err
		}
	} else {
		log.Debugf("Updating existsing BGPPeers: %s", peer.Name)
		if _, err := b.CalicoClient.BGPPeers().Update(context.Background(), peer, options.SetOptions{}); err != nil {
			return err
		}
	}

	return nil
}

func (b *BGPPeer) RemoveBGPPeer(peer *calicoApi.BGPPeer) error {
	_, err := b.CalicoClient.BGPPeers().Delete(context.Background(), peer.GetName(), options.DeleteOptions{})
	return err
}

func NewBGPPeer(calicoClient calicoClient.Interface) *BGPPeer {
	return &BGPPeer{
		CalicoClient: calicoClient,
	}
}
