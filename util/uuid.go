/*
 * @author: Haoyuan Liu
 * @date: 2020/12/2
 */

package util

import "github.com/google/uuid"

func NewUUID() uint32 {
	return uuid.New().ID()
}
